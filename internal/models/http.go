package models

import "time"

type RateByIDRequest struct {
	ID string `json:"id"`
}

type RateLatestRequest struct {
	Currency string `json:"currency"`
}

type RateUpdateRequest struct {
	Currency string `json:"currency"`
}

type UpdateRateResponse struct {
	ID string `json:"id"`
}

type RateResponse struct {
	Price      float32   `json:"price"`
	UpdateTime time.Time `json:"update_time"`
}
