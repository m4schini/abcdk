package s3

import (
	"fmt"
	"github.com/m4schini/abcdk/v3/errors"
	"net/url"
)

const (
	scheme = "s3"
)

type ConnectionString struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Secure    bool
	Bucket    string
}

func (c ConnectionString) String() string {
	u := url.URL{
		Scheme: scheme,
		Host:   c.Endpoint,
		User:   url.UserPassword(c.AccessKey, c.SecretKey),
		Path:   c.Bucket,
	}
	v := u.Query()
	v.Set("secure", fmt.Sprintf("%v", c.Secure))
	u.RawQuery = v.Encode()
	return u.String()
}

func ParseConn(connectionString string) (conn ConnectionString, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return conn, err
	}

	if u.Scheme != scheme {
		return conn, errors.ErrUnexpectedSchema
	}

	var isSet bool
	conn.Endpoint = u.Host
	conn.AccessKey = u.User.Username()
	conn.SecretKey, isSet = u.User.Password()
	if !isSet {
		return conn, errors.ErrConnStrIncomplete
	}
	conn.Secure = u.Query().Get("secure") != "false"
	if len(u.Path) > 0 {
		conn.Bucket = u.Path[1:]
	}

	return conn, nil
}
