package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App  App
		HTTP HTTP
		Log  Log
		PG   PG
		JWT  JWT
	}

	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
		ENV     string `env:"APP_ENVIRONMENT,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	JWT struct {
		Secret  string `env:"JWT_SECRET,required"`
		ExpDays int    `env:"JWT_EXP_DAYS,required"`
	}
)

func New() (*Config, error) {
	//if err := godotenv.Load(); err != nil {
	//	return nil, err
	//}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
