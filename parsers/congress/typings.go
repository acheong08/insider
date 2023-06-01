package congress

type Asset struct {
	Name         string `json:"name"`
	FilingStatus string `json:"filing_status"`
	Description  string `json:"description"`
	SubholdingOf string `json:"subholding_of"`
}
