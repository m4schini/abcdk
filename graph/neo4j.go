package graph

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"os"
	"time"
)

func NewFromEnv(ctx context.Context) (neo4j.SessionWithContext, context.CancelFunc, error) {
	uri := os.Getenv("NEO4J_URI")
	return New(ctx, uri)
}

func New(ctx context.Context, neo4jUri string) (neo4j.SessionWithContext, context.CancelFunc, error) {
	driver, err := neo4j.NewDriverWithContext(neo4jUri, neo4j.BasicAuth("neo4j", "northernlights", ""))
	if err != nil {
		return nil, func() {}, err
	}

	session := driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	return session, func() {
		closeCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		_ = session.Close(closeCtx)
		_ = driver.Close(closeCtx)
	}, nil
}
