package config

import "github.com/caarlos0/env/v6"

type Configuration struct {
	PGHost     string `env:"ENS_PG_HOST" envDefault:"127.0.0.1"`
	PGPort     string `env:"ENS_PG_PORT" envDefault:"5432"`
	PGDBName   string `env:"ENS_PG_DBNAME" envDefault:"learning"`
	PGUser     string `env:"ENS_PG_USER" envDefault:"postgres"`
	PGPassword string `env:"ENS_PG_PASSWORD" envDefault:"postgres"`
}

var config *Configuration

func init() {
	Init()
}

func Init() error {
	config = &Configuration{}
	err := env.Parse(config)
	return err
}

// Config() Get singleton instance
func Config() *Configuration {
	if config == nil || *config == (Configuration{}) {
		Init()
	}
	return config
}
