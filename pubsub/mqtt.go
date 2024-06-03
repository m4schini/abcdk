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

func FromConnStr(connStr string) (mqtt.Client, error) {
	conn, err := ParseConn(connStr)
	if err != nil {
		return nil, err
	}
	if conn.KeepAlive == 0 {
		conn.KeepAlive = 60 * time.Second
	}
	if conn.PingTimeout == 0 {
		conn.PingTimeout = 3 * time.Second
	}
	addr := fmt.Sprintf("%v://%v:%v", conn.Protocol, conn.Address, conn.Port)
	opts := mqtt.
		NewClientOptions().
		AddBroker(addr).
		SetClientID(conn.ClientId).
		SetKeepAlive(conn.KeepAlive).
		SetPingTimeout(conn.PingTimeout)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

func FromEnv() (mqtt.Client, error) {
	connstr := os.Getenv("PUBUSUB_URI")
	if connstr == "" {
		return nil, fmt.Errorf("PUBUSUB_URI cannot be empty")
	}
	return FromConnStr(connstr)
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
