package main

import (
	"context"
	"fmt"
	"github.com/m4schini/abcdk/pubsub"
)

func main() {
	s, err := pubsub.OpenSubscription(context.TODO(), "mqtt:///test-topic")
	if err != nil {
		panic(err)
	}

	for message := range s {
		fmt.Println(string(message.Payload))
	}
}
