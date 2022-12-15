package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"os"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 0, 20)

	buildHeading(m)
	buildInvoiceHeader(m)
	buildInvoicePositions(m)

	err := m.OutputFileAndClose("pdfs/test.pdf")
	if err != nil {
		fmt.Println("Couldn't safe PDF", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildInvoiceHeader(m pdf.Maroto) {
	m.Row(5, func() {
		m.Text("Blumenhaus Iris Martin | Bundesstr. 3 | 36764 Edingen", props.Text{
			Size:  9,
			Align: consts.Left,
		})
	})

	m.Row(12.7, func() {
		m.Col(50, func() {
			m.Text("Mahnung", props.Text{
				Size:  12,
				Align: consts.Left,
			})
		})
		m.ColSpace(8)
		m.Col(42, func() {
			m.Text("Rechnungsnummer:", props.Text{
				Size:  12,
				Align: consts.Left,
			})
			m.Text("000001", props.Text{
				Left:  45,
				Size:  12,
				Align: consts.Right,
			})

			m.Text("Inhaberin:", props.Text{
				Top:   7.5,
				Size:  12,
				Align: consts.Left,
			})

			m.Text("Iris Martin Fuhrländer", props.Text{
				Top:   7.5,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
		})
	})

	m.Row(27.3, func() {
		m.Col(50, func() {
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

		m.ColSpace(8)

		m.Col(42, func() {
			m.Text("Telefon: ", props.Text{
				Size:  12,
				Align: consts.Left,
			})
			m.Text("06449 488", props.Text{
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Fax", props.Text{
				Top:   5,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("06449 488", props.Text{
				Top:   5,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("E-Mail", props.Text{
				Top:   10,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("example@mail.com", props.Text{
				Top:   10,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
			m.Text("Datum", props.Text{
				Top:   17.5,
				Size:  12,
				Align: consts.Left,
			})
			m.Text("05.04.2001", props.Text{
				Top:   17.5,
				Left:  25,
				Size:  12,
				Align: consts.Right,
			})
		})
	})

	m.Row(15, func() {})

	m.Line(.1, props.Line{
		Color: color.NewBlack(),
		Style: consts.Solid,
		Width: 170,
	})
}

func buildHeading(m pdf.Maroto) {
	//Global header for every page
	m.RegisterHeader(func() {
		m.Row(45, func() {
			m.Text("Blumenhaus Iris Martin", props.Text{
				Top:             0,
				Left:            100,
				Right:           0,
				Size:            30,
				Align:           consts.Left,
				Extrapolate:     false,
				VerticalPadding: 0,
			})

			m.SetBackgroundColor(getTealColor())

			err := m.FileImage("images/logov2.jpg", props.Rect{
				Left:    160,
				Top:     0,
				Percent: 50,
				Center:  false,
			})

			if err != nil {
				fmt.Println("Image file was unable to load:", err)
			}
		})
	})
}

func buildInvoicePositions(m pdf.Maroto) {
	tableHeadings := []string{"Pos.", "Beschreibung", "Stk.", "Einzelpreis \n€", "Mwst \n%", "Gesamt \n€"}
	contents := [][]string{{"1", "Strauß (20)", "1", "20,00 ", "9,00", "9,80"}, {"2", "Strauß (15)", "2", "15,00 ", "7,00", "132,10"}}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{5, 65, 5, 11, 6, 8},
			Align:     []consts.Align{consts.Left, consts.Left, consts.Center, consts.Center, consts.Center, consts.Center},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{5, 65, 5, 11, 6, 8},
			Align:     []consts.Align{consts.Left, consts.Left, consts.Center, consts.Right, consts.Right, consts.Right},
		},
		HeaderContentSpace:     2,
		VerticalContentPadding: 1,
		Line:                   false,
		AlternatedBackground:   &lightPurpleColor,
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

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
