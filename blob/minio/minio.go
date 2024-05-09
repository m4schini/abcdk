package minio

import (
	"context"
	"fmt"
	abc "github.com/m4schini/abcdk/pkg/minio"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
)

type Bucket struct {
	client     *minio.Client
	bucketName string
	secure     bool
}

func OpenBucket(driverUrl *url.URL) (*Bucket, error) {
	c, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}
	b := new(Bucket)
	b.client = c
	b.bucketName = abc.Bucketname(driverUrl)
	_, b.secure = abc.Endpoint(driverUrl)
	return b, nil
}

func (b *Bucket) Upload(ctx context.Context, key string, reader io.Reader) (objectUrl string, err error) {
	info, err := b.client.PutObject(ctx, b.bucketName, key, reader, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/%v/%v", b.client.EndpointURL(), info.Bucket, info.Key), nil
}

func (b *Bucket) Download(ctx context.Context, key string) (io.Reader, error) {
	return b.client.GetObject(ctx, b.bucketName, key, minio.GetObjectOptions{})
}

func (b *Bucket) Delete(ctx context.Context, key string) error {
	return b.client.RemoveObject(ctx, b.bucketName, key, minio.RemoveObjectOptions{})
}
