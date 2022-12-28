package model

type Rechnungsposition struct {
	Positionsnummer    string
	Beschreibung       string
	Menge              string
	Einzelpreis        string
	Mehrwertsteuersatz string
	Gesammtpreis       string
}

func RechnungspositionenToContentStringSlice(rechnungspositionen []Rechnungsposition) [][]string {
	var rechnungspositionStringSlice [][]string
	for _, rechnungsposition := range rechnungspositionen {
		temp := []string{
			rechnungsposition.Positionsnummer,
			rechnungsposition.Beschreibung,
			rechnungsposition.Menge,
			rechnungsposition.Einzelpreis,
			rechnungsposition.Mehrwertsteuersatz,
			rechnungsposition.Gesammtpreis,
		}

		rechnungspositionStringSlice = append(rechnungspositionStringSlice, temp)
	}

	return rechnungspositionStringSlice
}
