package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

func SetLoggerDebug() {

}

func SetLoggerError() {

}

func NewFromEnv(clientId string) (mqtt.Client, error) {
	brokerAddr := os.Getenv("MQTT_BROKER")
	if brokerAddr == "" {
		return nil, fmt.Errorf("MQTT_BROKER cannot be empty")
	}
	return New(clientId, brokerAddr)
}

func New(clientId, broker string) (mqtt.Client, error) {
	opts := mqtt.
		NewClientOptions().
		AddBroker(broker).
		SetClientID(clientId).
		SetKeepAlive(60 * time.Second).
		SetPingTimeout(3 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}
