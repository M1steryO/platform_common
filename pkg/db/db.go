package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	DB() DB
	Close() error
}
type Query struct {
	Title string
	Query string
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dist interface{}, query Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dist interface{}, query Query, args ...interface{}) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
type Handler func(context.Context) error

type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

type Transactor interface {
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}
