package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DatabaseDSN string `env:"DATABASE_DSN"`
	Port        int    `env:"PORT" envDefault:"8080"`
	ClientCode  string `env:"CLIENT_CODE"`
	Username    string `env:"USERNAME"`
	Password    string `env:"PASSWORD"`
	Auth        string `env:"AUTH" envDefault:"admin"`
}

func ReadConfig() (Config, error) {
	cfgEnv := Config{}

	if err := env.Parse(&cfgEnv); err != nil {
		return cfgEnv, err
	}

	cfgFlag := Config{}

	flag.StringVar(&cfgFlag.DatabaseDSN, "d", cfgEnv.DatabaseDSN, "database DSN")
	flag.IntVar(&cfgFlag.Port, "p", cfgEnv.Port, "Port")
	flag.StringVar(&cfgFlag.ClientCode, "c", cfgEnv.ClientCode, "client code")
	flag.StringVar(&cfgFlag.Username, "u", cfgEnv.Username, "username")
	flag.StringVar(&cfgFlag.Password, "pass", cfgEnv.Password, "password")
	flag.StringVar(&cfgFlag.Auth, "a", cfgEnv.Auth, "auth")

	flag.Parse()

	return cfgFlag, nil
}
