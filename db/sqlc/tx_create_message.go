package db

import (
	"context"
)

type CreateMessageTxParams struct {
	CreateMessageParams
}

type CreateMessageTxResult struct {
	Message Message
}

func (store *SQLStore) CreateMessageTx(ctx context.Context, arg CreateMessageTxParams) (CreateMessageTxResult, error) {
	var result CreateMessageTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Message, err = q.CreateMessage(ctx, CreateMessageParams{
			From:    arg.From,
			Message: arg.Message,
		})
		return err
	})

	return result, err
}
