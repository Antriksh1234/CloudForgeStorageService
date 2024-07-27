package service

import "os"

type Service struct {
	StorageClient StorageClient
}

func (s *Service) StoreFileToCloud(file os.File, path string) error {
	return nil
}

func (s *Service) DownloadFileFromCloud(path string) (os.File, error) {
	return os.File{}, nil
}

func NewService() CloudForgeStorageService {
	return &Service{}
}
