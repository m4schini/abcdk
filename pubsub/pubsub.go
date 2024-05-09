package pubsub

import (
	"context"
	"github.com/m4schini/abcdk/errors"
	"github.com/m4schini/abcdk/internal/model"
	"net/url"
)

type Message struct {
	// Payload has the actual message content
	Payload []byte `json:"body"`
	// Metadata contains implementation dependent message metadata
	Metadata map[string]any `json:"metadata"`
}

// Topic represents a pubsub publisher
type Topic interface {
	// Send publish a message to this topic
	Send(ctx context.Context, message any) error
}

// OpenTopic opens a pubsub publisher
func OpenTopic(ctx context.Context, driverUrl string) (Topic, error) {
	d, err := url.Parse(driverUrl)
	if err != nil {
		return nil, err
	}
	switch d.Scheme {
	case model.SchemeMqtt:
		return openMqttTopic(ctx, d)
	case model.SchemeValkey, model.SchemeRedis:
		return openValkeyTopic(d)
	default:
		return nil, errors.ErrUnknownScheme
	}
}

// OpenSubscription opens a pubsub subscriber
func OpenSubscription(ctx context.Context, driverUrl string) (<-chan Message, error) {
	d, err := url.Parse(driverUrl)
	if err != nil {
		return nil, err
	}
	switch d.Scheme {
	case model.SchemeMqtt:
		return openMqttSubscription(ctx, d)
	case model.SchemeValkey, model.SchemeRedis:
		return openValkeySubscription(ctx, d)
	default:
		return nil, errors.ErrUnknownScheme
	}
}
