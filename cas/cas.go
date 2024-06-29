package cas

import (
	"context"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
)

type Client interface {
	Put(ctx context.Context, reader io.Reader, contentType string) (digest string, err error)
	Get(ctx context.Context, digest string) (object io.ReadCloser, contentType string, err error)
}

type client struct {
	s3     *minio.Client
	db     *mongo.Collection
	config ClientConfig
}

func New(s3 *minio.Client, db *mongo.Collection, config ClientConfig) Client {
	config = applyClientConfig(config)
	c := new(client)
	c.s3 = s3
	c.db = db
	c.config = config
	return c
}

func (c *client) Put(ctx context.Context, reader io.Reader, contentType string) (digest string, err error) {
	var (
		bucketName = c.config.BucketName
	)

	f, err := NewFSBuffer(reader)
	if err != nil {
		return "", err
	}

	r, err := f.Read()
	if err != nil {
		return "", err
	}
	defer r.Close()
	_, err = c.s3.PutObject(ctx, bucketName, f.Digest, r, f.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}
	return f.Digest, nil
}

func (c *client) Get(ctx context.Context, digest string) (object io.ReadCloser, contentType string, err error) {
	var (
		bucketName = c.config.BucketName
	)

	o, err := c.s3.GetObject(ctx, bucketName, digest, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}
	stat, err := o.Stat()
	if err != nil {
		return nil, "", err
	}
	return o, stat.ContentType, nil
}
