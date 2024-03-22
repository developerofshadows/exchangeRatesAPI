package config

import (
	"exchangeRatesAPI/internal/models"
	"gopkg.in/yaml.v3"
	"os"
)

func MustLoadConfig() *models.ServiceConfig {
	var cfgModel models.ServiceConfig

	f, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(f, &cfgModel)

	return &cfgModel
}
