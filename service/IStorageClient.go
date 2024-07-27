package service

import (
	"os"
)

type StorageClient interface {
	StoreFileToCloudStorage(file os.File, path string) error
	DownloadFileFromCloudStorage(path string) (os.File, error)
}
