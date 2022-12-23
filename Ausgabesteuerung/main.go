package Ausgabesteuerung

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
	"time"
)

func CreateInvoice() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	buildHeading(m)
	buildInvoiceHeader(m)
	buildInvoicePositions(m)
	buildFooter(m)

	err := m.OutputFileAndClose("Ausgabesteuerung/pdfs/test.pdf")
	if err != nil {
		fmt.Println("Couldn't safe PDF", err)
		os.Exit(1)
	}

	fmt.Printf("%v: PDF saved successfully", time.Now().Format("15:04:05"))
}

func buildHeading(m pdf.Maroto) {
	//Global header for every page
	m.RegisterHeader(func() {
		m.Row(42.5, func() {
			m.ColSpace(105)
			m.Col(65, func() {
				m.Text("Blumenhaus Iris Martin", props.Text{
					Top:             0,
					Right:           0,
					Size:            30,
					Align:           consts.Left,
					Extrapolate:     false,
					VerticalPadding: 0,
				})
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

func buildInvoiceHeader(m pdf.Maroto) {
	m.Row(5, func() {
		m.ColSpace(10)
		m.Col(170, func() {
			m.Text("Blumenhaus Iris Martin | Bundesstr. 3 | 36764 Edingen", props.Text{
				Size:  9,
				Align: consts.Left,
			})
		})

	})

	m.Row(12.7, func() {
		m.ColSpace(10)

		m.Col(85, func() {
			m.Text("Mahnung", props.Text{
				Size:  12,
				Align: consts.Left,
			})
		})

		m.ColSpace(20)

	})

	m.Row(27.3, func() {
		m.ColSpace(10)

		m.Col(85, func() {
			m.Text("Frau", props.Text{
				Size:  12,
				Align: consts.Left,
			})
			m.Text("Maren Dietrich", props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("Grünwetterstr. 24", props.Text{
				Top:   10,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("35745", props.Text{
				Top:   15,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("Herborn", props.Text{
				Top:   15,
				Left:  15,
				Size:  12,
				Align: consts.Left,
			})
		})

		m.ColSpace(20)

		m.Col(75, func() {
			m.Text("Test:", props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("000001", props.Text{
				Top:   5,
				Left:  45,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Rechnungsnummer: ", props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("000001", props.Text{
				Top:   5,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Rechnungsdatum:", props.Text{
				Top:   10,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("05.04.2021", props.Text{
				Top:   10,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Leistungsdatum:", props.Text{
				Top:   15,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("05.04.2022", props.Text{
				Top:   25,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
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

func buildInvoicePositions(m pdf.Maroto) {
	tableHeadings := []string{"Pos.", "Beschreibung", "Stk.", "Einzelpreis \n€", "Mwst \n%", "Gesamt \n€"}
	contents := [][]string{{"1", "Strauß (20)", "1", "20,00", "9,00", "9,80"}, {"2", "Strauß (15)", "2", "15,00", "7,00", "132,10"}}
	lightPurpleColor := getLightPurpleColor()
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
		},
		HeaderContentSpace:       2,
		VerticalContentPadding:   1,
		HorizontalContentPadding: 1,
		Line:                     false,
		AlternatedBackground:     &lightPurpleColor,
	})
}

func buildFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Line(.1, props.Line{
			Color: color.NewBlack(),
			Style: consts.Solid,
			Width: 170,
		})

		m.Row(2, func() {})

		m.Row(15, func() {
			m.Col(63, func() {
				m.Text("Blumenhaus Iris Martin", props.Text{
					Style: consts.Bold,
					Size:  8,
				})
				m.Text("Bundesstraße 3", props.Text{
					Top:  3,
					Size: 8,
				})
				m.Text("35764 Sinn", props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text("Telefon:", props.Text{
					Top:  9,
					Size: 8,
				})
				m.Text("06449-488:", props.Text{
					Top:  9,
					Left: 15,
					Size: 8,
				})
				m.Text("Fax:", props.Text{
					Top:  12,
					Size: 8,
				})
				m.Text("06449-6022", props.Text{
					Top:  12,
					Left: 15,
					Size: 8,
				})
				m.Text("Email:", props.Text{
					Top:  15,
					Size: 8,
				})
				m.Text("iris.mf.floristik@gmail.com", props.Text{
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
				m.Text("Sparkasse Dillenburg", props.Text{
					Top:  3,
					Size: 8,
				})
				m.Text("IBAN:", props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text("DE67 5165 0045 0000 0642 53", props.Text{
					Top:  6,
					Left: 10,
					Size: 8,
				})
				m.Text("BIC:", props.Text{
					Top:  9,
					Size: 8,
				})
				m.Text("HELADEF1DIL", props.Text{
					Top:  9,
					Left: 10,
					Size: 8,
				})
			})
			m.Col(63, func() {
				m.Text("Inhaberin", props.Text{
					Style: consts.Bold,
					Size:  8,
				})
				m.Text("Iris Martin-Fuhrländer", props.Text{
					Top:  3,
					Size: 8,
				})
				m.Text("Steuernr.:", props.Text{
					Top:  6,
					Size: 8,
				})
				m.Text("009/619/60117", props.Text{
					Top:  6,
					Left: 15,
					Size: 8,
				})
			})
		})
	})
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
