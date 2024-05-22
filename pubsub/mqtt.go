package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	abc "github.com/m4schini/abcdk/pkg/mqtt"
	"net/url"
	"time"
)

const (
	qos = 2
)

func openMqttTopic(ctx context.Context, driverUrl *url.URL) (Topic, error) {
	client, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}

	p := &mqttPublisher{
		topic:  abc.Topic(driverUrl),
		client: client,
	}
	return p, nil
}

func openMqttSubscription(ctx context.Context, driverUrl *url.URL) (<-chan Message, error) {
	client, err := abc.New(driverUrl)
	if err != nil {
		return nil, err
	}

	ch := make(chan Message)
	topic := abc.Topic(driverUrl)

	if token := client.Subscribe(topic, qos, func(client mqtt.Client, message mqtt.Message) {
		ch <- Message{
			Payload: message.Payload(),
			Metadata: map[string]any{
				"messageId": message.MessageID(),
				"retained":  message.Retained(),
				"qos":       message.Qos(),
			},
		}
	}); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	go func() {
		<-ctx.Done()
		client.Unsubscribe(topic).WaitTimeout(10 * time.Second)
		close(ch)
	}()

	return ch, nil
}

type mqttPublisher struct {
	topic  string
	client mqtt.Client
}

func (p *mqttPublisher) Send(ctx context.Context, message any) error {
	var data []byte
	switch x := message.(type) {
	case []byte:
		data = x
	case string:
		data = []byte(x)
	default:
		out, err := json.Marshal(&x)
		if err != nil {
			data = []byte(fmt.Sprintf("%v", x))
		} else {
			data = out
		}
	}

	token := p.client.Publish(p.topic, qos, false, data)
	token.Wait()
	return token.Error()
}
