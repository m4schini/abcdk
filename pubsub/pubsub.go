package pubsub

import "context"

type Message struct {
	Body     []byte
	Metadata map[string]string
}

type Topic interface {
	Send(ctx context.Context, message *Message) error
}

func OpenTopic(ctx context.Context, driverUrl string) (Topic, error) {

}

func OpenSubscription(ctx context.Context, driverUrl string) (<-chan *Message, error) {

}
