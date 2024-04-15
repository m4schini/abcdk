package sql

import (
	"context"
	"database/sql"
	"github.com/m4schini/abcdk/model"
)

func Open(ctx context.Context, driverUrl string) (*sql.DB, error) {
	return nil, model.ErrNotImplemented
}
