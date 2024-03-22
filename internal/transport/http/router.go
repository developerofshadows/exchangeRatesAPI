package httpT

import (
	"net/http"
)

// InitializeRouter Gets *http.ServeMux(router) as argument and adding paths and handlers for router
func InitializeRouter(mux *http.ServeMux) {
	rh := NewRateHandlers()

	// Adding update exchange rate path
	mux.Handle("PUT /api/v1/rate/update", http.HandlerFunc(rh.HandlerUpdateRate))

	// Adding get exchange rate by id path
	mux.Handle("POST /api/v1/rate/id", http.HandlerFunc(rh.HandlerGetRateByID))

	// Adding get latest exchange rate by currency path
	mux.Handle("POST /api/v1/rate/latest", http.HandlerFunc(rh.HandlerGetRateLatest))
}
