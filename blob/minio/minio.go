package minio

import (
	"context"
	"fmt"
	"github.com/m4schini/abcdk/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
)

var baseUrl = fmt.Sprintf("https://s3.%v", model.BaseUrl)

type Minio struct {
}

func NewConnstring(bucketName string, key, secret string) *url.URL {
	u := &url.URL{
		Scheme: "s3",
		Host:   bucketName,
	}
	u.Query().Set("key", key)
	u.Query().Set("secret", secret)
	return u
}

func ParseConnString(url *url.URL) (string, *credentials.Credentials, error) {
	bucketName := url.Host
	key := url.Query().Get("key")
	secret := url.Query().Get("secret")
	creds := credentials.NewStaticV4(key, secret, "")
	return bucketName, creds, nil
}

type Bucket struct {
	client     *minio.Client
	bucketName string
}

func OpenBucket(bucketName string, creds *credentials.Credentials) (*Bucket, error) {
	c, err := minio.New(baseUrl, &minio.Options{
		Creds:  creds,
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	b := new(Bucket)
	b.client = c
	return b, nil
}

func (b *Bucket) Upload(ctx context.Context, key string, reader io.Reader) (objectUrl string, err error) {
	info, err := b.client.PutObject(ctx, b.bucketName, key, reader, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v/%v/%v", baseUrl, info.Bucket, info.Key), nil
}

func (b *Bucket) Download(ctx context.Context, key string) (io.Reader, error) {
	return b.client.GetObject(ctx, b.bucketName, key, minio.GetObjectOptions{})
}

func (b *Bucket) Delete(ctx context.Context, key string) error {
	return b.client.RemoveObject(ctx, b.bucketName, key, minio.RemoveObjectOptions{})
}
