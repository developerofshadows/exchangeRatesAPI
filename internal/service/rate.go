package service

import (
	"exchangeRatesAPI/api/rate"
	"exchangeRatesAPI/internal/repository"
	"exchangeRatesAPI/pkg/logger"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type RateRepository interface {
	UpdateRate(fCurrency, sCurrency string, price float32, id string) error
	GetRateByUpdateIdentifier(id string) (float32, time.Time, error)
	GetLatestRate(fCurrency, sCurrency string) (float32, time.Time, error)
}

type RateExchangeAPI interface {
	GetLatestCurrencyRate(fCurrency, sCurrency string) (price float32)
}

type RateService struct {
	repo RateRepository
	api  RateExchangeAPI
}

func NewRateService() RateService {
	return RateService{
		repo: repository.NewRepository(),
		api:  rate.NewExchangeAPIInstance(),
	}
}

func (rs RateService) UpdateRate(fCurrency, sCurrency string) (id string, err error) {
	// TODO: Implement me
	id = uuid.New().String()
	go func() {
		logger.Logger.Debug(fmt.Sprintf("Updating price for: %s/%s", fCurrency, sCurrency))
		rs.repo.UpdateRate(fCurrency, sCurrency, rs.api.GetLatestCurrencyRate(fCurrency, sCurrency), id)
	}()

	return
}

func (rs RateService) GetLatestRate(fCurrency, sCurrency string) (price float32, updateTime time.Time, err error) {
	price, updateTime, err = rs.repo.GetLatestRate(fCurrency, sCurrency)

	return
}

func (rs RateService) GetRateByID(id string) (price float32, updateTime time.Time, err error) {
	price, updateTime, err = rs.repo.GetRateByUpdateIdentifier(id)

	return
}
