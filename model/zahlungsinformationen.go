package model

type Zahlungsinformationen struct {
	Bank               string `json:"bank"`
	Iban               string `json:"iban"`
	Bic                string `json:"bic"`
	Steuernummer       string `json:"steuernummer"`
	Umsatzsteuernummer string `json:"umsatzsteuernummer"`
}
