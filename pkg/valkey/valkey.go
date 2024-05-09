package valkey

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/m4schini/abcdk/internal/config"
	"github.com/valkey-io/valkey-go"
	"net/url"
)

var (
	cfg        = config.Config.Valkey
	clientName = fmt.Sprintf("abcdk_%v", uuid.New())
)

func Port(driverUrl *url.URL) string {
	port := driverUrl.Port()
	if port == "" {
		port = fmt.Sprintf("%v", cfg.Port)
	}
	return port
}

func Host(driverUrl *url.URL) string {
	host := driverUrl.Hostname()
	if host == "" {
		host = cfg.Host
	}
	return host
}

func Database(driverUrl *url.URL) string {
	database := driverUrl.Query().Get("database")
	if database == "" {
		database = "0"
	}
	return database
}

func New(driverUrl *url.URL) (valkey.Client, error) {
	opts := valkey.MustParseURL(fmt.Sprintf("redis://%v:%v/%v",
		Host(driverUrl),
		Port(driverUrl),
		Database(driverUrl)))
	opts.ClientName = clientName
	return valkey.NewClient(opts)
}
