package congress

type Asset struct {
	Name         string `json:"name"`
	Ticker       string `json:"ticker"`
	FilingStatus string `json:"filing_status"`
	Description  string `json:"description"`
	SubholdingOf string `json:"subholding_of"`
}

type Transaction struct {
	Owner            string `json:"owner"`
	Type             string `json:"type"`
	Date             string `json:"date"`
	NotificationDate string `json:"notification_date"`
	AmountRange      Range  `json:"amount_range"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}
