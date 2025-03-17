package config

import "os"

type Config struct {
	Port     string
	Timezone string
	SSLMode  string
	Db       *Db
}

func NewConfig() *Config {
	return &Config{
		Port:     os.Getenv("APP_PORT"),
		Timezone: os.Getenv("TIMEZONE"),
		SSLMode:  os.Getenv("SSL_MODE"),
		Db: &Db{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
		},
	}
}
