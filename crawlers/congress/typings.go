package congress

type DocumentEntry struct {
	Prefix        string `json:"prefix"`
	LastName      string `json:"last_name"`
	FirstName     string `json:"first_name"`
	Suffix        string `json:"suffix"`
	FillingType   string `json:"filling_type"`
	StateDistrict string `json:"state_district"`
	Year          string `json:"year"`
	FillingDate   string `json:"filling_date"`
	DocID         string `json:"doc_id"`
}
