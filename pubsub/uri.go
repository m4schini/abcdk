package pubsub

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	scheme              = "mqtt"
	keepaliveQueryKey   = "keepalive"
	pingTimeoutQueryKey = "pingTimeout"
)

type Protocol string

const (
	ProtocolTcp = "tcp"
	ProtocolSsl = "ssl"
	ProtocolWs  = "ws"
	ProtocolWss = "wss"
)

type ConnectionString struct {
	Address     string
	Protocol    Protocol
	ClientId    string
	KeepAlive   time.Duration
	PingTimeout time.Duration
}

func (c *ConnectionString) String() string {
	if c.Protocol == "" {
		c.Protocol = ProtocolTcp
	}
	u := &url.URL{
		Scheme: fmt.Sprintf("%v+%v", scheme, c.Protocol),
		Host:   c.Address,
		User:   url.User(c.ClientId),
	}
	q := u.Query()
	if c.KeepAlive != 0 {
		q.Set(keepaliveQueryKey, c.KeepAlive.String())
	}
	if c.PingTimeout != 0 {
		q.Set(pingTimeoutQueryKey, c.PingTimeout.String())
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func ParseConn(connectionString string) (conn ConnectionString, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return conn, err
	}

	var uScheme, uProtocol string
	uScheme, uProtocol, err = ParseScheme(u.Scheme)
	if err != nil {
		return conn, err
	}
	if uScheme != scheme {
		return conn, fmt.Errorf("invalid scheme")
	}
	if uProtocol == "" {
		uProtocol = ProtocolTcp
	}

	conn.Protocol = Protocol(uProtocol)
	conn.Address = u.Host
	conn.ClientId = u.User.Username()
	if u.Query().Has(keepaliveQueryKey) {
		d, err := time.ParseDuration(u.Query().Get(keepaliveQueryKey))
		if err == nil {
			conn.KeepAlive = d
		}
	}
	if u.Query().Has(pingTimeoutQueryKey) {
		d, err := time.ParseDuration(u.Query().Get(pingTimeoutQueryKey))
		if err == nil {
			conn.PingTimeout = d
		}
	}
	return conn, nil
}

func ParseScheme(schemeStr string) (scheme, protocol string, err error) {
	parts := strings.Split(schemeStr, "+")
	switch len(parts) {
	case 1:
		return parts[0], "", nil
	case 2:
		return parts[0], parts[1], nil
	default:
		return "", "", fmt.Errorf("invalid scheme")
	}
}
