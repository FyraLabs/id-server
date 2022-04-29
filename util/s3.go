package util

import (
	"github.com/fyralabs/id-server/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var S3Client *minio.Client

func InitializeS3() error {
	minioClient, err := minio.New(config.Environment.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Environment.S3AccessKey, config.Environment.S3SecretKey, ""),
		Secure: true,
	})

	if err != nil {
		return err
	}

	S3Client = minioClient

	return nil
}
