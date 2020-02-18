package env

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Env struct {
	DbDriver   string `env:"DB_DRIVER" envDefault:"mysql"`
	DbUser     string `env:"DB_USER" envDefault:"user"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DbHost     string `env:"DB_HOST" envDefault:"db"`
	DbPort     string `env:"DB_PORT" envDefault:"3306"`
	DbDatabase string `env:"DB_DATABASE" envDefault:"sample_db"`
}

// Load is load configs from a env file.
func Load() Env {
	cfg := Env{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("failed to load config %s", err)
	}
	return cfg
}
