package errors

import (
	"fmt"
	"github.com/m4schini/abcdk/v2/internal/consts"
)

var (
	ErrEnvVarIsMissing         = fmt.Errorf("env var is required")
	ErrDocstoreEnvVarIsMissing = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.DocstoreEnvVarName)
	ErrGraphEnvVarIsMissing    = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.GraphEnvVarName)
)
