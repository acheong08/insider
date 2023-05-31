package senate

var DRAW int = 0

type QueryParams struct {
	Draw               int      `json:"draw"`
	Columns            []column `json:"columns"`
	Order              []order  `json:"order"`
	Start              int      `json:"start"`
	Length             int      `json:"length"`
	SearchValue        string   `json:"search_value"`
	SearchRegex        bool     `json:"search_regex"`
	ReportTypes        []int    `json:"report_types"`
	FilterTypes        []int    `json:"filter_types"`
	SubmittedStartDate string   `json:"submitted_start_date"`
	SubmittedEndDate   string   `json:"submitted_end_date"`
	CandidateState     string   `json:"candidate_state"`
	SenatorState       string   `json:"senator_state"`
	OfficeID           string   `json:"office_id"`
	FirstName          string   `json:"first_name"`
	LastName           string   `json:"last_name"`
}

type column struct {
	Data        int    `json:"data"`
	Name        string `json:"name"`
	Searchable  bool   `json:"searchable"`
	Orderable   bool   `json:"orderable"`
	SearchValue string `json:"search_value"`
	SearchRegex bool   `json:"search_regex"`
}

type order struct {
	Column int    `json:"column"`
	Dir    string `json:"dir"`
}

func NewQueryParams(start, length int) QueryParams {
	DRAW += 1
	return QueryParams{
		Draw: DRAW,
		Columns: []column{
			// First name
			{
				Data:        0,
				Name:        "",
				Searchable:  true,
				Orderable:   true,
				SearchValue: "",
				SearchRegex: false,
			},
			// Last Name
			{
				Data:        1,
				Name:        "",
				Searchable:  true,
				Orderable:   true,
				SearchValue: "",
				SearchRegex: false,
			},
			// Office
			{
				Data:        2,
				Name:        "",
				Searchable:  true,
				Orderable:   true,
				SearchValue: "",
				SearchRegex: false,
			},
			// Report type
			{
				Data:        3,
				Name:        "",
				Searchable:  true,
				Orderable:   true,
				SearchValue: "",
				SearchRegex: false,
			},
			// Date filed
			{
				Data:        4,
				Name:        "",
				Searchable:  true,
				Orderable:   true,
				SearchValue: "",
				SearchRegex: false,
			},
		},
		Order: []order{
			{
				Column: 4,
				Dir:    "desc",
			},
		},
		Start:              start,
		Length:             length,
		SearchValue:        "",
		SearchRegex:        false,
		ReportTypes:        []int{11}, // Periodic transaction report
		FilterTypes:        []int{1},  // Filter for senators only
		SubmittedStartDate: "01/01/2012 00:00:00",
		SubmittedEndDate:   "",
		CandidateState:     "",
		SenatorState:       "",
		OfficeID:           "",
		FirstName:          "",
		LastName:           "",
	}
}

type MinimalQueryParams struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	FilterType          int    `json:"filter_type"`
	SenatorState        string `json:"senator_state"`
	ReportType          int    `json:"report_type"`
	SubmittedStartDate  string `json:"submitted_start_date"`
	SubmittedEndDate    string `json:"submitted_end_date"`
	Csrfmiddlewaretoken string `json:"csrfmiddlewaretoken"`
}

func NewMinimalQueryParams(qp QueryParams) MinimalQueryParams {
	return MinimalQueryParams{
		FilterType:          1,
		ReportType:          11,
		Csrfmiddlewaretoken: middleware_token,
	}
}

type Report struct {
	FirstName string
	LastName  string
	Office    string
	Ptr       string
	Date      string
}

type ReportResponse struct {
	Draw            int        `json:"draw"`
	RecordsTotal    int        `json:"recordsTotal"`
	RecordsFiltered int        `json:"recordsFiltered"`
	ReportData      [][]string `json:"data"`
	Result          string     `json:"result"`
}
