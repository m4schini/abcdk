package docstore

import (
	"context"
	"github.com/m4schini/abcdk/v3/errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"os"
	"time"
)

func FromEnv(ctx context.Context) (*mongo.Client, context.CancelFunc, error) {
	logger := zap.L().Named("abcdk").Named("docstore")
	uri := os.Getenv(consts.DocstoreEnvVarName)
	if uri == "" {
		err := errors.ErrDocstoreEnvVarIsMissing
		logger.Warn("failed to init docstore", zap.Error(err))
		return nil, func() {}, err
	}
	return FromConnStr(ctx, uri)
}

func FromConnStr(ctx context.Context, mongodbUri string) (*mongo.Client, context.CancelFunc, error) {
	logger := zap.L().Named("abcdk").Named("docstore")
	logger.Debug("connecting to mongodb")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		logger.Warn("failed to connect to mongodb", zap.Error(err))
		return nil, func() {}, err
	}

	logger.Debug("pinging mongodb")
	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		logger.Warn("failed to ping mondodb", zap.Error(err))
	}

	return client, func() {
		logger.Debug("disconnecting from mongodb")
		disconnectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		err := client.Disconnect(disconnectCtx)
		if err != nil {
			logger.Warn("failed to disconnect mongodb", zap.Error(err))
		}
	}, err
}
