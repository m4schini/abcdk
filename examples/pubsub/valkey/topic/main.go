package main

import (
	"context"
	"github.com/m4schini/abcdk/pubsub"
	"time"
)

func main() {
	t, err := pubsub.OpenTopic(context.TODO(), "valkey://localhost/example_test")
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
