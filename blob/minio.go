package blob

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

func FromConnStr(connStr string) (*minio.Client, error) {
	conn, err := ParseConn(connStr)
	if err != nil {
		return nil, err
	}
	return New(conn.Endpoint, conn.Key, conn.Secret)
}

func FromEnv() (*minio.Client, error) {
	connStr := os.Getenv("S3_URI")
	if connStr == "" {
		return nil, fmt.Errorf("S3_URI is required")
	}
	return FromConnStr(connStr)
}

func New(endpoint, key, secret string) (*minio.Client, error) {
	return minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(key, secret, ""),
		Secure: true,
	})
}
