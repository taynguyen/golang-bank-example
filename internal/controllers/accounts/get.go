package accounts

import (
	"context"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository/transactions"
)

func (i *impl) GetTransactions(ctx context.Context, userID int64) ([]models.Transaction, models.Pagination, error) {
	transactions, pagination, err := i.repo.Transactions().GetTransactions(ctx, transactions.GetTransactionsFilter{
		UserID: userID,
	})
	if err != nil {
		return nil, models.Pagination{}, err
	}
	return transactions, pagination, nil
}
