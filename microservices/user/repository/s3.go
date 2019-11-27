package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"io"
	"os"
)

type s3Repository struct {
}

func NewS3Repository() user.FileRepository {
	return s3Repository{}
}

func (r s3Repository) Store(file io.Reader) (location string, err error) {
	sess, err := session.NewSession()
	if err != nil {
		log.Error(err)
		return "", err
	}
	uploader := s3manager.NewUploader(sess)
	bucket := os.Getenv("AWS_BUCKET")
	uploadOutput, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(uuid.New().String()),
		Body:   file,
	})
	if err != nil {
		log.Error(err)
		return "", err
	}
	return uploadOutput.Location, nil
}

