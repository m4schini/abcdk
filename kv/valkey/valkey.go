package valkey

import (
	"context"
	abc "github.com/m4schini/abcdk/pkg/valkey"
	"github.com/valkey-io/valkey-go"
	"net/url"
)

func OpenKV(driverUrl *url.URL) (*valkeyKV, error) {
	client, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}
	return &valkeyKV{client: client}, nil
}

type valkeyKV struct {
	client valkey.Client
}

func (v *valkeyKV) Set(ctx context.Context, key string, value []byte) (err error) {
	cmd := v.client.B().
		Set().
		Key(key).
		Value(valkey.BinaryString(value)).Build()
	return v.client.Do(ctx, cmd).Error()
}

func (v *valkeyKV) Get(ctx context.Context, key string) (value []byte, err error) {
	cmd := v.client.B().Get().Key(key).Build()
	resp := v.client.Do(ctx, cmd)
	if err := resp.Error(); err != nil {
		return nil, err
	}
	return resp.AsBytes()
}
