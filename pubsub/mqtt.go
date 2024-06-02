package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/m4schini/abcdk/v2/internal/log"
	"os"
	"time"
)

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

func SetLoggerDebug(print func(v ...interface{}), printf func(format string, v ...interface{})) {
	mqtt.DEBUG = &log.MqttAdapter{
		PrintF:  print,
		PrintfF: printf,
	}
}

func SetLoggerError(print func(v ...interface{}), printf func(format string, v ...interface{})) {
	mqtt.ERROR = &log.MqttAdapter{
		PrintF:  print,
		PrintfF: printf,
	}
}

func SetLoggerCritical(print func(v ...interface{}), printf func(format string, v ...interface{})) {
	mqtt.CRITICAL = &log.MqttAdapter{
		PrintF:  print,
		PrintfF: printf,
	}
}

func SetLoggerWarn(print func(v ...interface{}), printf func(format string, v ...interface{})) {
	mqtt.WARN = &log.MqttAdapter{
		PrintF:  print,
		PrintfF: printf,
	}
}

func NewFromEnv(clientId string) (mqtt.Client, error) {
	brokerAddr := os.Getenv("MQTT_BROKER")
	if brokerAddr == "" {
		return nil, fmt.Errorf("MQTT_BROKER cannot be empty")
	}
	return New(clientId, brokerAddr)
}
