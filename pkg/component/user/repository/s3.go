package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/component/user"
	"github.com/go-park-mail-ru/2019_2_LeMMaS/pkg/logger"
	"github.com/google/uuid"
	"io"
	"os"
)

type s3Repository struct {
	logger logger.Logger
}

func NewS3Repository(logger logger.Logger) user.FileRepository {
	return s3Repository{logger}
}

func (r s3Repository) Store(file io.Reader) (location string, err error) {
	sess, err := session.NewSession()
	if err != nil {
		r.logger.Error(err)
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
		r.logger.Error(err)
		return "", err
	}
	return uploadOutput.Location, nil
}
