package config

import (
	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type EnvironmentType struct {
	JwtKey          string `env:"JWT_KEY,required=true"`
	DatabaseOptions string `env:"DATABASE_OPTIONS,required=true"`
}

var Environment EnvironmentType

func InitializeEnv() error {
	_ = godotenv.Load()
	_, err := env.UnmarshalFromEnviron(&Environment)
	if err != nil {
		return err
	}

	return nil
}
