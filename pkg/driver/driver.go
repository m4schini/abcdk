package driver

import "net/url"

type Url interface {
	AsDriverUrl() *url.URL
}
