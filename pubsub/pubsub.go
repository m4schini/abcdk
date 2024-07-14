package pubsub

import (
	"context"
	"errors"
	"fmt"
	mqtt2 "github.com/eclipse/paho.mqtt.golang"
	"github.com/m4schini/abcdk/v2/pubsub/mqtt"
	"github.com/m4schini/abcdk/v2/pubsub/uri"
	"github.com/m4schini/abcdk/v2/pubsub/valkey"
	vk "github.com/valkey-io/valkey-go"
	"os"
)

type Publisher interface {
	Publish(topic string, message []byte) error
}

type Subscriber interface {
	Subscribe(pattern string, handler func(topic string, message []byte)) error
}

type PubSub interface {
	Publisher
	Subscriber
}

func FromConnStr(connStr string) (PubSub, error) {
	conn, err := uri.ParseConn(connStr)
	if err != nil {
		return nil, err
	}
	switch conn.Scheme {
	case uri.SchemeMqtt:
		client, err := mqtt.FromConn(conn)
		if err != nil {
			return nil, err
		}
		return &mqttPubsub{client: client}, nil
	case uri.SchemeValkey:
		client, err := valkey.FromConn(conn)
		if err != nil {
			return nil, err
		}
		return &valkeyPubsub{client: client}, nil
	default:
		return nil, fmt.Errorf("unsupported scheme")
	}
}

func FromEnv() (PubSub, error) {
	connstr := os.Getenv("PUBSUB_URI")
	if connstr == "" {
		return nil, fmt.Errorf("PUBSUB_URI cannot be empty")
	}

	return FromConnStr(connstr)
}

type mqttPubsub struct {
	client mqtt2.Client
}

func (m *mqttPubsub) Publish(topic string, message []byte) error {
	t := m.client.Publish(topic, 2, false, message)
	t.Wait()
	return t.Error()
}

func (m *mqttPubsub) Subscribe(pattern string, handler func(topic string, message []byte)) error {
	t := m.client.Subscribe(pattern, 2, func(client mqtt2.Client, message mqtt2.Message) {
		handler(message.Topic(), message.Payload())
	})
	t.Wait()
	return t.Error()
}

type valkeyPubsub struct {
	client vk.Client
}

func (v *valkeyPubsub) Publish(topic string, message []byte) error {
	r := v.client.Do(context.Background(), v.client.B().
		Publish().
		Channel(topic).
		Message(string(message)).
		Build())
	return r.Error()
}

func (v *valkeyPubsub) Subscribe(pattern string, handler func(topic string, message []byte)) error {
	go func() {
		err := v.client.Receive(context.Background(), v.client.B().
			Psubscribe().
			Pattern(pattern).
			Build(),
			func(msg vk.PubSubMessage) {
				fmt.Println("received message")
				handler(msg.Channel, []byte(msg.Message))
			})
		if errors.Is(err, vk.ErrClosing) {
			return
		}
		if err != nil {
			panic(err)
		}
	}()
	return nil
}
