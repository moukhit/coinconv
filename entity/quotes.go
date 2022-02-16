package entity

type Quote struct {
	Code  string
	Price float64
}

type Quotes struct {
	List []Quote
}
