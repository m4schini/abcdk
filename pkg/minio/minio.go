package minio

import (
	"github.com/m4schini/abcdk/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"strings"
)

var cfg = config.Config.Minio

func Endpoint(driverUrl *url.URL) (endpoint string, secure bool) {
	endpoint = driverUrl.Host
	if endpoint == "" {
		endpoint = cfg.Endpoint
	}
	secureQuery := driverUrl.Query().Get("secure")
	if secureQuery == "" {
		secure = cfg.Secure
	} else {
		secure = secureQuery != "false"
	}
	return endpoint, secure
}

func Credentials(driverUrl *url.URL) *credentials.Credentials {
	key := driverUrl.Query().Get("key")
	secret := driverUrl.Query().Get("secret")
	return credentials.NewStaticV4(key, secret, "")
}

func Bucketname(driverUrl *url.URL) string {
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
