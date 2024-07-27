package config

import "os"

type Config struct {
	S3Bucket    string
	S3AWSRegion string
}

func InitializeConfigs() Config {
	return Config{
		S3Bucket:    os.Getenv("S3_BUCKET_NAME"),
		S3AWSRegion: os.Getenv("S3_AWS_REGION"),
	}
}
