package models

type ConfigModel struct {
	Env         string `env:"APP_ENV"`
	Port        string `env:"PORT"`
	AutoMigrate bool   `env:"AUTOMIGRATE"`
}
