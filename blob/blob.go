package blob

import (
	"context"
	"github.com/m4schini/abcdk/blob/minio"
	"github.com/m4schini/abcdk/errors"
	"github.com/m4schini/abcdk/internal/model"
	"io"
	"net/url"
)

type Bucket interface {
	Upload(ctx context.Context, key string, reader io.Reader) (objectUrl string, err error)
	Download(ctx context.Context, key string) (io.Reader, error)
	Delete(ctx context.Context, key string) error
}

func OpenBucket(driverUrl string) (Bucket, error) {
	d, err := url.Parse(driverUrl)
	if err != nil {
		return nil, err
	}
	switch d.Scheme {
	case model.SchemeS3:
		return minio.OpenBucket(d)
	default:
		return nil, errors.ErrUnknownScheme
	}
}
