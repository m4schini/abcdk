package minio

import (
	"fmt"
	"github.com/m4schini/abcdk/internal/config"
	"github.com/m4schini/abcdk/internal/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"strings"
)

const (
	queryKeyKey    = "key"
	queryKeySecret = "secret"
	queryKeySecure = "secure"
)

var cfg = config.Config.Minio

func Endpoint(driverUrl *url.URL) (endpoint string, secure bool) {
	endpoint = driverUrl.Host
	if endpoint == "" {
		endpoint = cfg.Endpoint
	}
	secureQuery := driverUrl.Query().Get(queryKeySecure)
	if secureQuery == "" {
		secure = cfg.Secure
	} else {
		secure = secureQuery != "false"
	}
	return endpoint, secure
}

func Credentials(driverUrl *url.URL) *credentials.Credentials {
	key := driverUrl.Query().Get(queryKeyKey)
	secret := driverUrl.Query().Get(queryKeySecret)
	return credentials.NewStaticV4(key, secret, "")
}

func BucketName(driverUrl *url.URL) string {
	return strings.ReplaceAll(driverUrl.Path, "/", "")
}

func New(driverUrl *url.URL) (*minio.Client, error) {
	endpoint, secure := Endpoint(driverUrl)
	creds := Credentials(driverUrl)
	return minio.New(endpoint, &minio.Options{
		Creds:  creds,
		Secure: secure,
	})
}

type Options struct {
	BucketName   string
	Endpoint     string
	Insecure     bool
	AccessKey    string
	AccessSecret string
}

func (o Options) AsDriverUrl() *url.URL {
	u := &url.URL{
		Scheme: model.SchemeS3,
		Host:   o.Endpoint,
		Path:   fmt.Sprintf("/%v", o.BucketName),
	}
	if o.AccessKey == "" || o.AccessSecret == "" {
		return nil
	}
	u.Query().Set(queryKeyKey, o.AccessKey)
	u.Query().Set(queryKeySecret, o.AccessSecret)
	if o.Insecure {
		u.Query().Set(queryKeySecure, "false")
	}
	return u
}
