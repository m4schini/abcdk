package pubsub

import (
	"github.com/m4schini/abcdk/v2/pubsub/uri"
	"testing"
	"time"
)

func TestValkeyPubsub(t *testing.T) {
	conn := uri.ConnectionString{
		Address: "127.0.0.1",
		Port:    6379,
		Scheme:  uri.SchemeValkey,
	}

	t.Log("parsing connection string")
	p, err := FromConnStr(conn.String())
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("publishing")
	err = p.Publish("test/123", []byte("example message"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("published")

	t.Log("subscribing")
	err = p.Subscribe("test/*", func(topic string, message []byte) {
		t.Log(topic, string(message))
	})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("subscribed")

	time.Sleep(3 * time.Second)

	t.Log("publishing")
	err = p.Publish("test/123", []byte("example message"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log("published")

	time.Sleep(3 * time.Second)
}
