package kv

import (
	"github.com/m4schini/abcdk/v3/errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"github.com/valkey-io/valkey-go"
	"go.uber.org/zap"
	"os"
)

func FromEnv() (valkey.Client, error) {
	logger := zap.L().Named("abcdk").Named("kv")
	uri := os.Getenv(consts.KVEnvVarName)
	if uri == "" {
		err := errors.ErrKVEnvVarIsMissing
		logger.Warn("failed to init valkey client", zap.Error(err))
		return nil, err
	}

	return FromConnStr(uri)
}

func FromConnStr(valkeyUri string) (valkey.Client, error) {
	logger := zap.L().Named("abcdk").Named("kv")
	conn, err := ParseConn(valkeyUri)
	if err != nil {
		return nil, err
	}

	logger.Debug("creating valkey client")
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{conn.Address},
	})
	if err != nil {
		logger.Warn("failed to create valkey client", zap.Error(err))
	}
	return client, nil
}
