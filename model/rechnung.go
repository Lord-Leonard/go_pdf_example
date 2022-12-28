package model

import "github.com/Lord-Leonard/go_pdf_example/date"

type Rechnung struct {
	Kunde  Person
	Klient Firma

	Wichtigkeit     string
	Rechnungsnummer string
	RechnungsDatum  date.Date
	LeistungsDatum  date.Date
	LeistungVon     date.Date
	LeistungBis     date.Date
	Positionen      []Rechnungsposition
	Netto           string
	Mehrwertsteuer  string
	Gesamtpreis     string
}
