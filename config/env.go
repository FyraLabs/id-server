package config

import (
	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type EnvironmentType struct {
	JwtKey            string `env:"JWT_KEY,required=true"`
	DatabaseOptions   string `env:"DATABASE_OPTIONS,required=true"`
	SendGridKey       string `env:"SENDGRID_KEY,required=true"`
	S3AccessKey       string `env:"S3_ACCESS_KEY,required=true"`
	S3SecretKey       string `env:"S3_SECRET_KEY,required=true"`
	S3Endpoint        string `env:"S3_ENDPOINT,required=true"`
	S3Bucket          string `env:"S3_BUCKET,required=true"`
	S3AvatarURLPrefix string `env:"S3_AVATAR_URL_PREFIX,required=true"`
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
