package entity

// Quote is a conversion quote for particular currency
type Quote struct {
	Code  string
	Price float64
}

// Quotes contains list of quotes
type Quotes struct {
	List []Quote
}
