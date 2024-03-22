package models

type ExchangeAPIResponse struct {
	Rates struct {
		EUR float32 `json:"EUR"`
		MXN float32 `json:"MXN"`
		USD float32 `json:"USD"`
	} `json:"rates"`
}

var Currencies = map[string]struct{}{
	"USD": {},
	"EUR": {},
	"MXN": {},
}
