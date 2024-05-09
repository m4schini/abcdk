package kv

import (
	"context"
	"fmt"
	"github.com/m4schini/abcdk/errors"
	"github.com/m4schini/abcdk/internal/model"
	"github.com/m4schini/abcdk/kv/valkey"
	"net/url"
)

type KV interface {
	Set(ctx context.Context, key string, value []byte) (err error)
	Get(ctx context.Context, key string) (value []byte, err error)
}

func OpenKV(ctx context.Context, driverUrl string) (KV, error) {
	d, err := url.Parse(driverUrl)
	if err != nil {
		return nil, err
	}
	switch d.Scheme {
	case model.SchemeValkey, model.SchemeRedis:
		host := d.Host
		db := d.Path
		fmt.Println(host, db)
		return valkey.OpenKV(d)
	default:
		return nil, errors.ErrUnknownScheme
	}
}
