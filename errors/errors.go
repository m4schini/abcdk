package errors

import (
	"fmt"
	"github.com/m4schini/abcdk/v3/internal/consts"
)

var (
	ErrEnvVarIsMissing         = fmt.Errorf("env var is required")
	ErrDocstoreEnvVarIsMissing = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.DocstoreEnvVarName)
	ErrGraphEnvVarIsMissing    = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.GraphEnvVarName)
)
