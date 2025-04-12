package docstore

import (
	"github.com/m4schini/abcdk/v3/internal/consts"
	"net/url"
)

type ConnectionString struct {
	Server   string
	Username string
	Password string
}

func (c *ConnectionString) String() string {
	u := url.URL{
		Scheme: consts.SchemeDocstore,
		User:   url.UserPassword(c.Username, c.Password),
		Host:   c.Server,
	}
	return u.String()
}
