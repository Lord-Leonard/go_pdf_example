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
	buildInvoicePositions(m)

	err := m.OutputFileAndClose("pdfs/test.pdf")
	if err != nil {
		fmt.Println("Couldn't safe PDF", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
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
	tableHeadings := []string{"Pos.", "Beschreibung", "Menge", "Einzelpreis €", "Gesamt €"}
	contents := [][]string{{"1", "Strauß (20)", "1", "20,00", "20,00"}, {"2", "Strauß (15)", "2", "15,00", "30,00"}}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Positionen", props.Text{
				Top:    2,
				Size:   12,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 6, 1, 2, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 6, 1, 2, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
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
