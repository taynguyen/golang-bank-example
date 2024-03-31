package transactions

import (
	"context"
	"gin-boilerplate/internal/models"
)

type GetTransactionsFilter struct {
	UserID int64
}

func (i *impl) GetTransactions(ctx context.Context, filter GetTransactionsFilter) ([]models.Transaction, models.Pagination, error) {
	// transactions, pagination, err := i.db.
	// if err != nil {
	// 	return nil, models.Pagination{}, err
	// }
	// return transactions, pagination, nil

	return nil, models.Pagination{}, nil
}
