package kv

import (
	"github.com/m4schini/abcdk/v3/errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"net/url"
)

type ConnectionString struct {
	Address string
}

func (c ConnectionString) String() string {
	u := url.URL{
		Scheme: consts.SchemeKv,
		Host:   c.Address,
	}
	return u.String()
}

func ParseConn(connectionString string) (conn ConnectionString, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return conn, err
	}

	if u.Scheme != consts.SchemeKv {
		return conn, errors.ErrUnexpectedSchema
	}

	conn.Address = u.Host
	return conn, nil
}
