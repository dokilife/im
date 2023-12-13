package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Querier
	AcquireConn(ctx context.Context) (*pgxpool.Conn, error)
	CreateMessageTx(ctx context.Context, arg CreateMessageTxParams) (CreateMessageTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (s *SQLStore) AcquireConn(ctx context.Context) (*pgxpool.Conn, error) {
	return s.connPool.Acquire(ctx)
}
