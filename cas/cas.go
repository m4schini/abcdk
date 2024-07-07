package cas

import (
	"context"
	"errors"
	"github.com/m4schini/abcdk/v2/cas/internal/buffer"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
)

type Stat struct {
	Digest      string              `bson:"digest"`
	ContentType string              `bson:"contentType"`
	MetaData    map[string][]string `bson:"metadata"`
}

type Storage interface {
	Put(ctx context.Context, reader io.Reader, contentType string) (digest string, err error)
	Get(ctx context.Context, digest string) (object io.ReadCloser, contentType string, err error)
	Stat(ctx context.Context, digest string) (stat Stat, err error)
}

type Client struct {
	S3     *minio.Client
	DB     *mongo.Collection
	config ClientConfig
}

func New(s3 *minio.Client, db *mongo.Client, config ClientConfig) *Client {
	config = applyClientConfig(config)
	c := new(Client)
	c.S3 = s3
	c.DB = db.Database(config.DatabaseName).Collection(config.CollectionName)
	c.config = config
	return c
}

func (c *Client) Put(ctx context.Context, reader io.Reader, contentType string) (digest string, err error) {
	var (
		bucketName = c.config.BucketName
	)

	f, err := buffer.NewFSBuffer(reader)
	if err != nil {
		return "", err
	}
	defer f.Delete()

	r, err := f.Read()
	if err != nil {
		return f.Digest, err
	}
	defer r.Close()
	_, err = c.S3.PutObject(ctx, bucketName, f.Digest, r, f.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return f.Digest, err
	}

	opts := options.Update().SetUpsert(true)
	_, err = c.DB.UpdateOne(ctx, bson.D{{"digest", f.Digest}}, bson.D{{"$set", Stat{
		Digest:      f.Digest,
		ContentType: contentType,
	}}}, opts)
	return f.Digest, err
}

func (c *Client) Get(ctx context.Context, digest string) (object io.ReadCloser, contentType string, err error) {
	var (
		bucketName = c.config.BucketName
	)

	o, err := c.S3.GetObject(ctx, bucketName, digest, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}
	stat, err := o.Stat()
	if err != nil {
		return nil, "", err
	}
	return o, stat.ContentType, nil
}

func (c *Client) Stat(ctx context.Context, digest string) (stat Stat, err error) {
	err = c.Data(ctx, digest, &stat)
	return stat, err
}

func (c *Client) Data(ctx context.Context, digest string, v any) (err error) {
	r := c.DB.FindOne(ctx, bson.D{{"digest", digest}})
	err = r.Err()
	if err != nil {
		return err
	}

	err = r.Decode(v)
	return err
}

func GetWithStat(ctx context.Context, storage Storage, digest string) (object io.ReadCloser, stat Stat, err error) {
	errCh := make(chan error)
	go func() {
		var err error
		stat, err = storage.Stat(ctx, digest)
		errCh <- err
	}()

	object, _, err = storage.Get(ctx, digest)
	err = errors.Join(err, <-errCh)
	return object, stat, err
}
