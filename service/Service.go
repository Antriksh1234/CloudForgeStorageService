package service

import (
	"os"

	"github.com/CloudForgeMonitoringService/config"
)

type Service struct {
	StorageClient StorageClient
	Config        config.Config
}

func (s *Service) StoreFileToCloud(file os.File, path string) error {
	return nil
}

func (s *Service) DownloadFileFromCloud(path string) (os.File, error) {
	return os.File{}, nil
}

func NewService() (CloudForgeStorageService, error) {
	S3Client, err := NewStorageClient()

	return &Service{
		StorageClient: S3Client,
		Config:        config.InitializeConfigs(),
	}, err
}
