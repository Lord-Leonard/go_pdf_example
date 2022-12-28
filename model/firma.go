package model

type Firma struct {
	Name                  string
	Telefon               string
	Fax                   string
	Email                 string
	Steuernummer          string
	Umsatzsteuernummer    string
	Geschäftsführer       Person
	Adresse               Adresse
	Zahlungsinformationen Zahlungsinformationen
}
