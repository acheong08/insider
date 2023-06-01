package senate

type Transaction struct {
	Date        string `json:"date"`
	Owner       string `json:"owner"`
	Ticker      string `json:"ticker"`
	Asset       Asset  `json:"asset"`
	Type        string `json:"type"`
	AmountRange Range  `json:"amount_range"`
	Comment     string `json:"comment"`
}

type Asset struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	OptionType  string  `json:"option_type"`
	StrikePrice float64 `json:"strike_price"`
	Expires     string  `json:"expires"`
}

type Range struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
