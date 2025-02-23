package config

import "os"

type HTTPConfig struct {
	Host       string
	Port       string
	ExposePort string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host:       os.Getenv("APP_HOST"),
		Port:       os.Getenv("APP_PORT"),
		ExposePort: os.Getenv("EXPOSE_PORT"),
	}
}
