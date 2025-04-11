package errors

import (
	"errors"
	"testing"
)

func TestErrorIsEnvarIsMissing(t *testing.T) {
	errs := []error{
		ErrDocstoreEnvVarIsMissing,
		ErrGraphEnvVarIsMissing,
	}
	for _, err := range errs {
		t.Log(err.Error())
		if !errors.Is(err, ErrEnvVarIsMissing) {
			t.Errorf("expected error (%v) to be ErrEnvVarIsMissing", err)
			t.Fail()
		}
	}
}
