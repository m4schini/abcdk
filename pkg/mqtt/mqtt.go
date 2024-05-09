package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/m4schini/abcdk/internal/config"
	"net/url"
)

const DefaultProtocol = "tcp"

var (
	cfg      = config.Config.Mqtt
	clientId = fmt.Sprintf("abcdk_%v", uuid.New())
)

func toPort(scheme string) int {
	port, ok := cfg.Listeners[scheme]
	if !ok {
		return -1
	}
	return port
}

func Broker(driverUrl *url.URL) string {
	protocol := driverUrl.Query().Get("protocol")
	if protocol == "" {
		protocol = DefaultProtocol
	}
	hostname := driverUrl.Hostname()
	if hostname == "" {
		hostname = cfg.Host
	}
	port := driverUrl.Port()
	if port == "" {
		port = fmt.Sprintf("%v", toPort(protocol))
	}
	return fmt.Sprintf("%v://%v:%v", protocol, hostname, port)
}

func Topic(driverUrl *url.URL) string {
	if len(driverUrl.Path) < 2 {
		panic("invalid topic")
	}
	return driverUrl.Path[1:]
}

func New(driverUrl *url.URL) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(Broker(driverUrl))
	opts.SetClientID(clientId)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}
