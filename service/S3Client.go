package service

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
}

func (s *S3Client) StoreFileToCloudStorage(file os.File, path string) error {
	return nil
}

func (s *S3Client) DownloadFileFromCloudStorage(path string) (os.File, error) {
	return os.File{}, nil
}

func NewStorageClient() (StorageClient, error) {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion(os.Getenv("S3_AWS_REGION")))

	if err != nil {
		log.Println("Error while creating the config for S3!", err.Error())
		return &S3Client{}, err
	}

	s3Client := s3.NewFromConfig(config)

	return &S3Client{
		Client: s3Client,
	}, nil
}
