package httpT

import (
	"encoding/json"
	"net/http"
)

var (
	RespondErrorBadID       = "Bad ID"
	RespondErrorBadCurrency = "Invalid Currency"
)

func RespondWithErrorMessage(msg string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	_ = json.NewEncoder(w).Encode(msg)
}

func RespondWithInternalServerError(err error, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	_ = json.NewEncoder(w).Encode(err.Error())
}
