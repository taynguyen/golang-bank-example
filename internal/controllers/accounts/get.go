package accounts

import (
	"context"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository/transactions"
)

func (i *impl) GetTransactions(ctx context.Context, filter transactions.GetTransactionsFilter) ([]models.Transaction, models.Pagination, error) {
	transactions, pagination, err := i.repo.Transactions().GetTransactions(ctx, filter)
	if err != nil {
		return nil, models.Pagination{}, err
	}
	return transactions, pagination, nil
}
