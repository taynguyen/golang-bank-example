package transactions

import (
	"context"
	"gin-boilerplate/internal/models"

	"github.com/pkg/errors"
)

type GetTransactionsFilter struct {
	models.Pagination

	UserID *uint
}

func (i *impl) GetTransactions(ctx context.Context, filter GetTransactionsFilter) ([]models.Transaction, models.Pagination, error) {
	stmt := i.db.Model(&models.Transaction{}).Joins("JOIN accounts ON transactions.account_id = accounts.id")

	if filter.UserID != nil {
		stmt = stmt.Where("accounts.user_id = ?", *filter.UserID)
	}

	paging := filter.Pagination

	err := stmt.Count(&paging.Total).Error
	if err != nil {
		return nil, paging, errors.WithStack(err)
	}

	var transactions []models.Transaction
	err = stmt.Limit(paging.Limit).Offset(paging.Offset).Find(&transactions).Error
	if err != nil {
		return nil, paging, errors.WithStack(err)
	}

	return transactions, paging, nil
}
