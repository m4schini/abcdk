package graph

import (
	"fmt"
	"github.com/m4schini/abcdk/v3/errors"
	"github.com/m4schini/abcdk/v3/internal/consts"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/url"
	"strconv"
)

const (
	accessModeQueryKey = "accessMode"
)

type ConnectionString struct {
	Host       string
	Username   string
	Password   string
	AccessMode *neo4j.AccessMode
}

func (c *ConnectionString) String() string {
	u := url.URL{
		Scheme: consts.SchemeGraph,
		User:   url.UserPassword(c.Username, c.Password),
		Host:   c.Host,
	}
	if c.AccessMode != nil {
		u.Query().Set(accessModeQueryKey, fmt.Sprintf("%v", *c.AccessMode))
	}

	return u.String()
}

func ParseConn(connectionString string) (conn ConnectionString, err error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		return conn, err
	}

	if u.Scheme != consts.SchemeGraph {
		return conn, errors.ErrUnexpectedSchema
	}

	conn.Host = u.Host
	accessMode := neo4j.AccessModeRead
	conn.AccessMode = &accessMode
	if u.User != nil {
		conn.Username = u.User.Username()
		conn.Password, _ = u.User.Password()
	}
	if u.Query().Has(accessModeQueryKey) {
		mode, err := strconv.ParseInt(u.Query().Get(accessModeQueryKey), 10, 64)
		if err == nil {
			m := neo4j.AccessMode(mode)
			conn.AccessMode = &m
		}
	}
	return conn, nil
}
