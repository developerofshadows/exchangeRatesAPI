package models

type ServiceConfig struct {
	HTTP            `yaml:"http"`
	Logger          `yaml:"logger"`
	Database        `yaml:"database"`
	ServiceSpecific `yaml:"service_spec"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Logger struct {
	LoggerLevel string `yaml:"logger_level"`
}

type HTTP struct {
	Port int `yaml:"port"`
}

type ServiceSpecific struct {
	BaseCurrency string `yaml:"base_currency"`
}
