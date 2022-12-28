package Ausgabesteuerung

import (
	"fmt"
	"github.com/Lord-Leonard/go_pdf_example/model"
	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"net/http"
	"os"
	"time"
)

func RechnungErstellen(rechnung model.Rechnung, c *gin.Context) string {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	buildHeading(m, rechnung.Klient)
	buildInvoiceHeader(m, rechnung.Klient, rechnung.Kunde, rechnung)
	buildInvoicePositions(m, rechnung.Positionen)
	buildInvoiceTotal(m, rechnung.Netto, rechnung.Mehrwertsteuer, rechnung.Gesamtpreis)
	buildFooter(m, rechnung.Klient)

	pdfDir := fmt.Sprintf("Kunden/%s/Rechnungen", rechnung.Kunde.Name)
	pdfPath := fmt.Sprintf("Kunden/%s/Rechnungen/%s.pdf", rechnung.Kunde.Name, rechnung.Rechnungsnummer)

	err := os.MkdirAll(pdfDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = m.OutputFileAndClose(pdfPath)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	return pdfPath
}

func buildHeading(m pdf.Maroto, klient model.Firma) {
	//Global header for every page
	m.RegisterHeader(func() {
		m.Row(42.5, func() {
			m.ColSpace(105)
			m.Col(65, func() {
				m.Text(
					klient.Name,
					props.Text{
						Top:             0,
						Right:           0,
						Size:            30,
						Align:           consts.Left,
						Extrapolate:     false,
						VerticalPadding: 0,
					},
				)
			})

			m.Col(20, func() {
				err := m.FileImage("Ausgabesteuerung/images/logov2.jpg", props.Rect{
					Top:    2,
					Center: false,
				})

				if err != nil {
					fmt.Println("Image file was unable to load:", err)
				}
			})

		})
	})
}

func buildInvoiceHeader(m pdf.Maroto, klient model.Firma, kunde model.Person, rechnung model.Rechnung) {
	m.Row(5, func() {
		m.ColSpace(10)
		m.Col(170, func() {
			m.Text(fmt.Sprintf("%s | %s %s | %s %s",
				klient.Name,
				klient.Adresse.Straße,
				klient.Adresse.Hausnummer,
				klient.Adresse.Postleitzahl,
				klient.Adresse.Stadt,
			),
				props.Text{
					Size:  9,
					Align: consts.Left,
				})
		})

	})

	m.Row(12.7, func() {
		m.ColSpace(10)

		m.Col(85, func() {
			m.Text(rechnung.Wichtigkeit, props.Text{
				Size:  12,
				Align: consts.Left,
			})
		})

		m.ColSpace(20)

	})

	m.Row(27.3, func() {
		m.ColSpace(10)

		m.Col(85, func() {
			m.Text(kunde.Andrede, props.Text{
				Size:  12,
				Align: consts.Left,
			})

			m.Text(kunde.Name, props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Left,
			})

			m.Text(fmt.Sprintf("%s %s",
				kunde.Adresse.Straße,
				kunde.Adresse.Hausnummer,
			),
				props.Text{
					Top:   10,
					Size:  12,
					Align: consts.Left,
				})

			m.Text(kunde.Adresse.Postleitzahl, props.Text{
				Top:   15,
				Size:  12,
				Align: consts.Left,
			})

			m.Text(kunde.Adresse.Stadt, props.Text{
				Top:   15,
				Left:  15,
				Size:  12,
				Align: consts.Left,
			})
		})

		m.ColSpace(20)

		m.Col(75, func() {
			m.Text("Rechnungsnummer: ", props.Text{
				Size:  12,
				Align: consts.Left,
			})
			m.Text(rechnung.Rechnungsnummer, props.Text{
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Rechnungsdatum:", props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Left,
			})
			m.Text(time.Time(rechnung.RechnungsDatum).Format("02.04.2006"), props.Text{
				Top:   5,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			if !time.Time(rechnung.LeistungsDatum).IsZero() {
				m.Text("Leistungsdatum:", props.Text{
					Top:   10,
					Size:  12,
					Align: consts.Left,
				})
				m.Text(time.Time(rechnung.LeistungsDatum).Format("02.04.2006"), props.Text{
					Top:   10,
					Left:  25,
					Size:  12,
					Align: consts.Right,
				})
			} else {
				m.Text("Leistungszeitraum:", props.Text{
					Top:   10,
					Size:  12,
					Align: consts.Left,
				})
				m.Text(time.Time(rechnung.LeistungVon).Format("02.04.2006"), props.Text{
					Top:   10,
					Left:  25,
					Size:  12,
					Align: consts.Right,
				})
				m.Text(fmt.Sprintf("- %s",
					time.Time(rechnung.LeistungBis).Format("02.04.2006")),
					props.Text{
						Top:   15,
						Left:  25,
						Size:  12,
						Align: consts.Right,
					})
			}
		})
	})

	m.Row(10, func() {})

	m.Line(.1, props.Line{
		Color: color.NewBlack(),
		Style: consts.Solid,
		Width: 170,
	})

	m.Row(2.5, func() {})
}

func buildInvoicePositions(m pdf.Maroto, positionen []model.Rechnungsposition) {
	tableHeadings := []string{"Pos.", "Beschreibung", "Stk.", "Einzelpreis", "Mwst", "Gesamt"}
	contents := model.RechnungspositionenToContentStringSlice(positionen)
	gridSice := []uint{10, 115, 10, 25, 13, 17}

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: gridSice,
			Align:     []consts.Align{consts.Left, consts.Left, consts.Center, consts.Center, consts.Center, consts.Center},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: gridSice,
			Align:     []consts.Align{consts.Center, consts.Left, consts.Center, consts.Right, consts.Right, consts.Right},
			Unit:      []consts.Unit{consts.None, consts.None, consts.None, consts.Euro, consts.Percent, consts.Euro},
		},
		HeaderContentSpace:       2,
		VerticalContentPadding:   1,
		HorizontalContentPadding: 3,
		Line:                     true,
	})

	m.Row(5, func() {})

}

func buildInvoiceTotal(m pdf.Maroto, netto string, mehrwertsteuer string, gesamtpreis string) {
	m.Row(12, func() {
		m.Col(110, func() {
			m.Text("Der Gesamtbetrag ist nach erhalt dieser Rechnung ohne Abzüge zu zahlen ")
		})

		m.ColSpace(10)

		m.Col(70, func() {
			m.Text("Nettobetrag", props.Text{
				Size: 11,
			})
			m.Text(netto, props.Text{
				Align: consts.Right,
				Size:  11,
				Left:  30,
				Unit:  consts.Euro,
			})

			m.Text("zzgl. Mwst.", props.Text{
				Size: 11,
				Top:  6,
			})
			m.Text(mehrwertsteuer, props.Text{
				Align: consts.Right,
				Size:  11,
				Top:   6,
				Left:  30,
				Unit:  consts.Euro,
			})

			m.Text("Gesamtbetrag", props.Text{
				Size:  11,
				Top:   12,
				Style: consts.Bold,
			})
			m.Text(gesamtpreis, props.Text{
				Align: consts.Right,
				Size:  11,
				Top:   12,
				Left:  30,
				Style: consts.Bold,
				Unit:  consts.Euro,
			})
		})
	})
}

func buildFooter(m pdf.Maroto, klient model.Firma) {
	m.RegisterFooter(func() {
		m.Line(.1, props.Line{
			Color: color.NewBlack(),
			Style: consts.Solid,
			Width: 150,
		})

		m.Row(2, func() {})

		m.Row(15, func() {

			m.ColSpace(10)

			m.Col(58, func() {
				m.Text(klient.Name, props.Text{
					Style: consts.Bold,
					Size:  8,
				})
				m.Text(fmt.Sprintf("%s %s", klient.Adresse.Straße, klient.Adresse.Hausnummer), props.Text{
					Top:  3,
					Size: 8,
				})
				m.Text(fmt.Sprintf("%s %s", klient.Adresse.Postleitzahl, klient.Adresse.Stadt), props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text("Telefon:", props.Text{
					Top:  9,
					Size: 8,
				})
				m.Text(klient.Telefon, props.Text{
					Top:  9,
					Left: 15,
					Size: 8,
				})
				m.Text("Fax:", props.Text{
					Top:  12,
					Size: 8,
				})
				m.Text(klient.Fax, props.Text{
					Top:  12,
					Left: 15,
					Size: 8,
				})
				m.Text("Email:", props.Text{
					Top:  15,
					Size: 8,
				})
				m.Text(klient.Email, props.Text{
					Top:  15,
					Left: 15,
					Size: 8,
				})
			})

			m.Col(64, func() {
				m.Text("Bankverbindung", props.Text{
					Style: consts.Bold,
					Size:  8,
				})
				m.Text(klient.Zahlungsinformationen.Bank, props.Text{
					Top:  3,
					Size: 8,
				})
				m.Text("IBAN:", props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text(klient.Zahlungsinformationen.Iban, props.Text{
					Top:  6,
					Left: 10,
					Size: 8,
				})
				m.Text("BIC:", props.Text{
					Top:  9,
					Size: 8,
				})
				m.Text(klient.Zahlungsinformationen.Bic, props.Text{
					Top:  9,
					Left: 10,
					Size: 8,
				})
			})

			m.Col(58, func() {
				m.Text("Inhaberin", props.Text{
					Style: consts.Bold,
					Size:  8,
				})
				m.Text(klient.Geschäftsführer.Name, props.Text{
					Top:  3,
					Size: 8,
				})

				m.Text("Steuernr.:", props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text(klient.Steuernummer, props.Text{
					Top:  6,
					Left: 25,
					Size: 8,
				})

				m.Text("Umsatzsteuernr.:", props.Text{
					Top:  9,
					Size: 8,
				})
				m.Text(klient.Umsatzsteuernummer, props.Text{
					Top:  9,
					Left: 25,
					Size: 8,
				})
			})
		})
	})
}
