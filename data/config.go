package data

import (
	"encuestify/models"

	"github.com/caarlos0/env/v6"
)

var Config models.ConfigModel = GetConfig()

func GetConfig() models.ConfigModel {
	cfg := models.ConfigModel{}

	err := env.Parse(&cfg)

	if err != nil {
		panic(1)
	}

	return cfg
}
