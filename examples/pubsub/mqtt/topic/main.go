package main

import (
	"context"
	"github.com/m4schini/abcdk/pubsub"
	"time"
)

func main() {
	t, err := pubsub.OpenTopic(context.TODO(), "mqtt:///test-topic")
	if err != nil {
		panic(err)
	}

	for {
		err = t.Send(context.TODO(), "test_message")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}
