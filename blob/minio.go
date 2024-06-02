package blob

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

func NewFromEnv() (*minio.Client, error) {
	endpoint := os.Getenv("S3_ENDPOINT")
	key := os.Getenv("S3_KEY")
	secret := os.Getenv("S3_SECRET")
	return New(endpoint, key, secret)
}

func New(endpoint, key, secret string) (*minio.Client, error) {
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(key, secret, ""),
		Secure: true,
	})
}
