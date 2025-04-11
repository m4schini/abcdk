package graph

import (
	"context"
	"github.com/m4schini/abcdk/v2/errors"
	"github.com/m4schini/abcdk/v2/internal/consts"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.uber.org/zap"
	"os"
	"time"
)

func FromEnv(ctx context.Context) (neo4j.SessionWithContext, context.CancelFunc, error) {
	logger := zap.L().Named("abcdk").Named("graph")
	uri := os.Getenv(consts.GraphEnvVarName)
	if uri == "" {
		err := errors.ErrGraphEnvVarIsMissing
		logger.Warn("failed to init graph", zap.Error(err))
		return nil, func() {}, err
	}
	return FromConnStr(ctx, uri)
}

func FromConnStr(ctx context.Context, neo4jUri string) (neo4j.SessionWithContext, context.CancelFunc, error) {
	logger := zap.L().Named("abcdk").Named("graph")
	logger.Debug("connecting to neo4j")
	driver, err := neo4j.NewDriverWithContext(neo4jUri, neo4j.BasicAuth("neo4j", "northernlights", ""))
	if err != nil {
		logger.Warn("failed to connect to neo4j", zap.Error(err))
		return nil, func() {}, err
	}

	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	return session, func() {
		logger.Debug("closing neo4j session")
		closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		err := session.Close(closeCtx)
		if err != nil {
			logger.Warn("failed to close neo4j session", zap.Error(err))
		}
		err = driver.Close(closeCtx)
		if err != nil {
			logger.Warn("failed to close neo4j session", zap.Error(err))
		}
	}, nil
}
