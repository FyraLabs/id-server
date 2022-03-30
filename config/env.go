package config

import (
	env "github.com/Netflix/go-env"
)

type Environment struct {
	JwtKey string `env:"JWT_KEY,required=true"`
}

var environment Environment

func InitializeEnv() error {
	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return err
	}

	return err
}
