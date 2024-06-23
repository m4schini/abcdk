package blob

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

type Bucket interface {
	Put(ctx context.Context, key string, reader io.Reader, size int64, options minio.PutObjectOptions) (err error)
	Get(ctx context.Context, key string, options minio.GetObjectOptions) (object io.ReadCloser, err error)
}

type minioBucket struct {
	s3         *minio.Client
	bucketName string
}

func NewBucket(client *minio.Client, bucketName string) (*minioBucket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("bucket doesn't exist")
	}

	b := new(minioBucket)
	b.s3 = client
	b.bucketName = bucketName
	return b, nil
}

func (m *minioBucket) Put(ctx context.Context, key string, reader io.Reader, size int64, options minio.PutObjectOptions) (err error) {
	_, err = m.s3.PutObject(ctx, m.bucketName, key, reader, size, options)
	return err
}

func (m *minioBucket) Get(ctx context.Context, key string, options minio.GetObjectOptions) (object io.ReadCloser, err error) {
	return m.s3.GetObject(ctx, m.bucketName, key, options)
}
