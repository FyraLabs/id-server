package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fyralabs/id-server/config"
)

var UploadClient *s3manager.Uploader
var S3Client *s3.S3

func InitializeS3() error {
	session, err := session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("auto"),
		Endpoint:         aws.String(config.Environment.S3Endpoint),
		Credentials:      credentials.NewStaticCredentials(config.Environment.S3AccessKey, config.Environment.S3SecretKey, ""),
	})

	if err != nil {
		return err
	}

	S3Client = s3.New(session)
	UploadClient = s3manager.NewUploader(session)

	return nil
}
