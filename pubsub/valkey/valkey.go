package valkey

import (
	"fmt"
	"github.com/m4schini/abcdk/v2/pubsub/uri"
	"github.com/valkey-io/valkey-go"
)

func FromConn(connectionString uri.ConnectionString) (valkey.Client, error) {
	if connectionString.Scheme != uri.SchemeValkey {
		return nil, fmt.Errorf("unsupported scheme")
	}

	return valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{fmt.Sprintf("%v:%v", connectionString.Address, connectionString.Port)},
	})
}
