package errors

import (
	"fmt"
	"github.com/m4schini/abcdk/v3/internal/consts"
)

var (
	ErrEnvVarIsMissing         = fmt.Errorf("env var is required")
	ErrDocstoreEnvVarIsMissing = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.DocstoreEnvVarName)
	ErrGraphEnvVarIsMissing    = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.GraphEnvVarName)
	ErrS3EnvVarIsMissing       = fmt.Errorf("%w: %v", ErrEnvVarIsMissing, consts.S3EnvVarName)

	ErrUnexpectedSchema  = fmt.Errorf("unexpected schema")
	ErrConnStrIncomplete = fmt.Errorf("conn str is incomplete")
)
