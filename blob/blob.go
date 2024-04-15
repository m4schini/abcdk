package blob

import (
	"context"
	"github.com/m4schini/abcdk/blob/minio"
	"github.com/m4schini/abcdk/model"
	"io"
	"net/url"
)

const (
	SchemeS3  = "s3"
	SchemeMem = "mem"
)

type Bucket interface {
	Upload(ctx context.Context, key string, reader io.Reader) (objectUrl string, err error)
	Download(ctx context.Context, key string) (io.Reader, error)
	Delete(ctx context.Context, key string) error
}

func OpenBucket(ctx context.Context, uri string) (Bucket, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case SchemeS3:
		bucketName, creds, _ := minio.ParseConnString(u)
		return minio.OpenBucket(bucketName, creds)
	default:
		return nil, model.ErrUnknownScheme
	}
}
