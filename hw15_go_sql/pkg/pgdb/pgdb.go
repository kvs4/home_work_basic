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

func New(ctx context.Context, DSN string) error {
	dbc, err := pgxpool.New(ctx, DSN)
	if err != nil {
		return errors.Errorf("failed to connect to DB: %v", err)
	}

	if err := dbc.Ping(ctx); err != nil {
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
