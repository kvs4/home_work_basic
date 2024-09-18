package pgdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Connection struct {
	dbc *pgxpool.Pool
}

var DB Connection

func New(ctx context.Context, dsn string) error {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return errors.Errorf("failed to connect to DB: %v", err)
	}

	err = dbc.Ping(ctx)
	if err != nil {
		return errors.Errorf("failed to ping DB: %v", err)
	}

	DB = Connection{
		dbc: dbc,
	}

	return nil
}

func (p *Connection) Conn() *pgxpool.Pool {
	return p.dbc
}
