package docstore

import "net/url"

type ConnectionString struct {
	Server   string
	Username string
	Password string
}

func (c *ConnectionString) String() string {
	u := url.URL{
		Scheme: "mongodb",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   c.Server,
	}
	return u.String()
}
