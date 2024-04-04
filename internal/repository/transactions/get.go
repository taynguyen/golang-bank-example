package transactions

import (
	"context"
	"gin-boilerplate/internal/models"

	"github.com/pkg/errors"
)

type GetTransactionsFilter struct {
	models.Pagination

	UserID    *uint
	AccountID *uint
}

func (i *impl) GetTransactions(ctx context.Context, filter GetTransactionsFilter) ([]models.Transaction, models.Pagination, error) {
	stmt := i.db.WithContext(ctx).Model(&models.Transaction{}).Joins("JOIN accounts ON transactions.account_id = accounts.id")

	if filter.UserID != nil {
		stmt = stmt.Where("accounts.user_id = ?", *filter.UserID)
	}

	if filter.AccountID != nil {
		stmt = stmt.Where("transactions.account_id = ?", *filter.AccountID)
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
