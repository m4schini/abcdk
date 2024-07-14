package valkey

import (
	"github.com/m4schini/abcdk/v2/pubsub/uri"
	"golang.org/x/net/context"
	"testing"
)

func TestFromConn(t *testing.T) {
	client, err := FromConn(uri.ConnectionString{
		Address: "127.0.0.1",
		Port:    6379,
		Scheme:  uri.SchemeValkey,
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	client.Do(context.TODO(), client.B().Ping().Build())
}
