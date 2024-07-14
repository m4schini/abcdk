package uri

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	keepaliveQueryKey   = "keepalive"
	pingTimeoutQueryKey = "pingTimeout"
)

type Protocol string

const (
	SchemeMqtt              = "mqtt"
	SchemeValkey            = "valkey"
	ProtocolTcp    Protocol = "tcp"
	defaultPortTcp          = 1883
	ProtocolSsl    Protocol = "ssl"
	defaultPortSsl          = 8883
	ProtocolWs     Protocol = "ws"
	defaultPortWs           = 8083
	ProtocolWss    Protocol = "wss"
	defaultPortWss          = 8084
)

type ConnectionString struct {
	Address     string
	Port        int
	Scheme      string
	Protocol    Protocol
	ClientId    string
	KeepAlive   time.Duration
	PingTimeout time.Duration
}

func (c *ConnectionString) String() string {
	var scheme string
	switch {
	case c.Scheme == SchemeMqtt && c.Protocol == "":
		scheme = fmt.Sprintf("%v+%v", c.Scheme, ProtocolTcp)
		break
	case c.Protocol == "":
		scheme = c.Scheme
		break
	default:
		scheme = fmt.Sprintf("%v+%v", c.Scheme, c.Protocol)
		break
	}
	if c.Port == 0 {
		c.Port = Port(c.Protocol)
	}

	u := &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%v:%v", c.Address, c.Port),
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
	if uScheme != SchemeMqtt && uScheme != SchemeValkey {
		return conn, fmt.Errorf("invalid scheme")
	}
	conn.Scheme = uScheme
	if uProtocol == "" {
		uProtocol = string(ProtocolTcp)
	}
	conn.Protocol = Protocol(uProtocol)

	hostParts := strings.Split(u.Host, ":")
	switch len(hostParts) {
	case 1:
		conn.Address = hostParts[0]
		conn.Port = Port(conn.Protocol)
		break
	case 2:
		conn.Address = hostParts[0]
		port, err := strconv.ParseInt(hostParts[1], 10, 32)
		if err != nil {
			return conn, err
		}
		conn.Port = int(port)
	default:
		return conn, fmt.Errorf("invalid host")
	}

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

func Port(protocol Protocol) int {
	switch protocol {
	case ProtocolTcp:
		return defaultPortTcp
	case ProtocolSsl:
		return defaultPortSsl
	case ProtocolWs:
		return defaultPortWs
	case ProtocolWss:
		return defaultPortWss
	}
	return 0
}
