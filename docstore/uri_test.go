package docstore

import "testing"

func TestConnectionString_String(t *testing.T) {
	conn := ConnectionString{
		Server:   "172.30.0.4:27017",
		Username: "m4schini",
		Password: "password",
	}
	t.Log(conn.String())
}
