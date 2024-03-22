package fawazahmed0

import (
	"encoding/json"
	"exchangeRatesAPI/internal/models"
	"exchangeRatesAPI/pkg/logger"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ExchangeAPI struct{}

func (ea ExchangeAPI) GetLatestCurrencyRate(fCurrency, sCurrency string) (price float32) {
	resp, err := http.Get(fmt.Sprintf("https://api.frankfurter.app/latest?amount=1&from=%s&to=%s", fCurrency, sCurrency))
	if err != nil {
		logger.Logger.Error("Cannot get to the Exchange API")
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result models.ExchangeAPIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		logger.Logger.Error("Cannot unmarshal Exchange API answer")
		os.Exit(0)
	}

	switch sCurrency {
	case "EUR":
		return result.Rates.EUR
	case "MXN":
		return result.Rates.MXN
	case "USD":
		return result.Rates.USD
	default:
		logger.Logger.Error("Unknown currency type")
		return
	}
}
