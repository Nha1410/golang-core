package config

import (
	"os"
	"strconv"

	"github.com/caarlos0/env"
)

type DBConfig struct {
	User            string `env:"DB_USER" envDefault:""`
	Password        string `env:"DB_PASSWORD" envDefault:""`
	Driver          string `env:"DB_DRIVER" envDefault:""`
	Name            string `env:"DB_NAME" envDefault:""`
	Host            string `env:"DB_HOST" envDefault:""`
	Port            string `env:"DB_PORT" envDefault:""`
	Timezone        string `env:"DB_TIMEZONE" envDefault:""`
	ConnMaxIdleTime int    `env:"DB_CONN_MAX_IDLE_TIME" envDefault:"0"`
	MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS" envDefault:"0"`
	MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS" envDefault:"0"`
	LogSkip         bool
}

func LoadDBConfig() DBConfig {
	dbCfg := DBConfig{}
	if err := env.Parse(&dbCfg); err != nil {
		panic(err)
	}

	var logSkip bool
	if uss := os.Getenv("DB_LOG_SKIP"); uss != "" {
		v, err := strconv.ParseBool(uss)
		if err != nil {
			panic(err)
		}
		logSkip = v
	}

	dbCfg.LogSkip = logSkip

	return dbCfg
}
