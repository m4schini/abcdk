package pubsub

import "context"

func ExampleOpenTopic() {
	// mqtt://[<hostname>][:<port>]/<topic>[?protocol=tcp|ssl|ws|wss]
	_, _ = OpenTopic(context.TODO(), "mqtt:///mqtt_topic")

	// keyval://[<hostname>][:<port>]/<topic>[?database=<databaseId>]
	_, _ = OpenTopic(context.TODO(), "keyval:///keyval_topic")
}

func ExampleOpenSubscription() {
	// mqtt://[<hostname>][:<port>]/<topic>[?protocol=tcp|ssl|ws|wss]
	_, _ = OpenSubscription(context.TODO(), "mqtt:///mqtt_topic")

	// keyval://[<hostname>][:<port>]/<topic>[?database=<databaseId>]
	_, _ = OpenSubscription(context.TODO(), "keyval:///keyval_topic")
}
