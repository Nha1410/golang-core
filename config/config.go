package config

type Config struct {
	DB             DBConfig
	HTTP           HTTPConfig
	UseSwaggerSpec bool
	BaseUrl        string
}

func NewConfig() *Config {
	return &Config{
		DB:             LoadDBConfig(),
		HTTP:           LoadHTTPConfig(),
		UseSwaggerSpec: true,
		BaseUrl:        "/api/v1",
	}
}
