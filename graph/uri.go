package graph

import (
	"fmt"
	"net/url"
)

const (
	scheme = "neo4j"
)

type ConnectionString struct {
	Host     string
	Username string
	Password string
}

func (c *ConnectionString) String() string {
	u := url.URL{
		Scheme: scheme,
		User:   url.UserPassword(c.Username, c.Password),
		Host:   c.Host,
	}
	return u.String()
}

func ParseConn(connectionString string) (conn ConnectionString, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return conn, err
	}

	if u.Scheme != scheme {
		return conn, fmt.Errorf("invalid scheme")
	}

	conn.Host = u.Host
	if u.User != nil {
		conn.Username = u.User.Username()
		conn.Password, _ = u.User.Password()
	}
	return conn, nil
}
