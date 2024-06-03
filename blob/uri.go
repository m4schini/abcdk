package blob

import (
	"fmt"
	"net/url"
)

const (
	scheme = "s3"
)

type ConnectionString struct {
	Endpoint   string
	Key        string
	Secret     string
	BucketName string
}

func (c *ConnectionString) String() string {
	u := url.URL{
		Scheme: scheme,
	}
	if c.Endpoint != "" {
		u.Host = c.Endpoint
	}
	if c.Secret != "" || c.Key != "" {
		u.User = url.UserPassword(c.Key, c.Secret)
	}
	if c.BucketName != "" {
		u.Path = c.BucketName
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

	conn.Endpoint = u.Host
	if u.User != nil {
		conn.Key = u.User.Username()
		conn.Secret, _ = u.User.Password()
	}

	if len(u.Path) > 0 {
		conn.BucketName = u.Path[1:]
	}

	return conn, nil
}
