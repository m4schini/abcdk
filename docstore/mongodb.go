package docstore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

func FromEnv(ctx context.Context) (*mongo.Client, context.CancelFunc, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, func() {}, fmt.Errorf("MONGODB_URI required")
	}
	return FromConnStr(ctx, uri)
}

func FromConnStr(ctx context.Context, mongodbUri string) (*mongo.Client, context.CancelFunc, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, func() {}, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err = client.Ping(pingCtx, readpref.Primary())
	return client, func() {
		disconnectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		_ = client.Disconnect(disconnectCtx)
	}, err
}
