package errors

import "fmt"

var ErrUnknownScheme = fmt.Errorf("unknown scheme")
var ErrNotImplemented = fmt.Errorf("not implemented")
