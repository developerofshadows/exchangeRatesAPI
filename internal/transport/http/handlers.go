package httpT

import (
	"encoding/json"
	"exchangeRatesAPI/internal/models"
	"exchangeRatesAPI/internal/service"
	"net/http"
	"strings"
	"time"
)

type RateService interface {
	UpdateRate(fCurrency, sCurrency string) (string, error)
	GetLatestRate(fCurrency, sCurrency string) (float32, time.Time, error)
	GetRateByID(id string) (float32, time.Time, error)
}

type RateHandlers struct {
	s RateService
}

func NewRateHandlers() RateHandlers {
	return RateHandlers{
		s: service.NewRateService(),
	}
}

// @Summary UpdateRate
// @tags update
// @Description Update exchange rate for specified currencies
// @ID update-rate
// @Accept json
// @Produce json
// @Param input body models.RateUpdateRequest true "Currency pair"
// @Success 200 {object} models.UpdateRateResponse
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /api/v1/rate/update [put]
func (rh RateHandlers) HandlerUpdateRate(w http.ResponseWriter, r *http.Request) {
	var err error
	var inp models.RateUpdateRequest
	var resp models.UpdateRateResponse

	if err = json.NewDecoder(r.Body).Decode(&inp); err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	curs := strings.Split(inp.Currency, "/")

	if inp.Currency == "" || len(curs) != 2 {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	if _, ok := models.Currencies[curs[0]]; !ok {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	if _, ok := models.Currencies[curs[1]]; !ok {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	resp.ID, err = rh.s.UpdateRate(curs[0], curs[1])

	if err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	RespondWithData(resp, w)
}

// @Summary GetLatestRate
// @tags get
// @Description Get latest exchange rate for specified currencies
// @ID latest-rate
// @Accept json
// @Produce json
// @Param input body models.RateLatestRequest true "Currency pair"
// @Success 200 {float} float 1
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /api/v1/rate/latest [post]
func (rh RateHandlers) HandlerGetRateLatest(w http.ResponseWriter, r *http.Request) {
	var err error
	var inp models.RateLatestRequest
	var resp models.RateResponse

	if err = json.NewDecoder(r.Body).Decode(&inp); err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	if inp.Currency == "" {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	curs := strings.Split(inp.Currency, "/")

	if inp.Currency == "" || len(curs) != 2 {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	if _, ok := models.Currencies[curs[0]]; !ok {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	if _, ok := models.Currencies[curs[1]]; !ok {
		RespondWithErrorMessage(RespondErrorBadCurrency, w, r)
		return
	}

	resp.Price, resp.UpdateTime, err = rh.s.GetLatestRate(curs[0], curs[1])

	if err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	RespondWithData(resp, w)
}

// @Summary GetRateByID
// @tags get
// @Description Get exchange rate by ID for specified currencies
// @ID id-rate
// @Accept json
// @Produce json
// @Param input body models.RateByIDRequest true "ID"
// @Success 200 {object} models.RateResponse
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /api/v1/rate/id [post]
func (rh RateHandlers) HandlerGetRateByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var inp models.RateByIDRequest
	var resp models.RateResponse

	if err = json.NewDecoder(r.Body).Decode(&inp); err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	if inp.ID == "" {
		RespondWithErrorMessage(RespondErrorBadID, w, r)
		return
	}

	resp.Price, resp.UpdateTime, err = rh.s.GetRateByID(inp.ID)

	if err != nil {
		RespondWithInternalServerError(err, w, r)
		return
	}

	RespondWithData(resp, w)
}

func RespondWithData(data any, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(data)
}
