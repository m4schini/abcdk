package pubsub

import (
	"context"
	"github.com/m4schini/abcdk/model"
)

type Message struct {
	Body     []byte
	Metadata map[string]string
}

type Topic interface {
	Send(ctx context.Context, message *Message) error
}

func OpenTopic(ctx context.Context, driverUrl string) (Topic, error) {
	return nil, model.ErrNotImplemented
}

func OpenSubscription(ctx context.Context, driverUrl string) (<-chan *Message, error) {
	return nil, model.ErrNotImplemented
}
