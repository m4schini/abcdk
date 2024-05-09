package pubsub

import (
	"context"
	"fmt"
	abc "github.com/m4schini/abcdk/pkg/valkey"
	"github.com/valkey-io/valkey-go"
	"net/url"
)

type valkeyPublisher struct {
	topic  string
	client valkey.Client
}

func openValkeyTopic(driverUrl *url.URL) (*valkeyPublisher, error) {
	client, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}
	return &valkeyPublisher{
		topic:  driverUrl.Path,
		client: client,
	}, nil
}

func (v *valkeyPublisher) Send(ctx context.Context, message any) error {
	cmd := v.client.B().Publish().Channel(v.topic).Message(fmt.Sprintf("%v", message)).Build()
	return v.client.Do(ctx, cmd).Error()
}

func openValkeySubscription(ctx context.Context, driverUrl *url.URL) (<-chan Message, error) {
	client, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}
	topic := driverUrl.Path
	ch := make(chan Message)
	go func() {
		err = client.Receive(ctx, client.B().Subscribe().Channel(topic).Build(), func(msg valkey.PubSubMessage) {
			ch <- Message{
				Payload: []byte(msg.Message),
				Metadata: map[string]any{
					"pattern": msg.Pattern,
					"channel": msg.Channel,
				},
			}
		})
		close(ch)
	}()
	return ch, nil
}
