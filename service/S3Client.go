package service

import (
	"os"

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

func NewStorageClient() StorageClient {
	return &S3Client{}
}
