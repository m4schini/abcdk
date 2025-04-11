package s3

import (
	"github.com/m4schini/abcdk/v3/errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"os"
)

func FromEnv() (client *minio.Client, bucketName string, err error) {
	logger := zap.L().Named("abcdk").Named("s3")
	uri := os.Getenv(consts.S3EnvVarName)
	if uri == "" {
		err := errors.ErrS3EnvVarIsMissing
		logger.Warn("failed to init s3 client", zap.Error(err))
		return nil, "", err
	}
	return FromConnStr(uri)
}

func FromConnStr(minioUri string) (client *minio.Client, bucketName string, err error) {
	logger := zap.L().Named("abcdk").Named("s3")
	conn, err := ParseConn(minioUri)
	if err != nil {
		return nil, "", err
	}

	logger.Debug("creating minio client")
	client, err = minio.New(conn.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conn.AccessKey, conn.SecretKey, ""),
		Secure: conn.Secure,
	})
	if err != nil {
		logger.Warn("failed to create minio client", zap.Error(err))
	}
	return client, conn.Bucket, err
}
