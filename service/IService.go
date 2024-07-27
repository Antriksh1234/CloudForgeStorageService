package service

import "os"

type CloudForgeStorageService interface {
	StoreFileToCloud(file os.File, path string) error
	DownloadFileFromCloud(path string) (os.File, error)
}
