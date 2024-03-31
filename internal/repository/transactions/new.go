package transactions

import (
	"context"
	"gin-boilerplate/internal/models"

	"gorm.io/gorm"
)

type ITransactionRepo interface {
	GetTransactions(ctx context.Context, filter GetTransactionsFilter) ([]models.Transaction, models.Pagination, error)
}

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITransactionRepo {
	return &impl{
		db: db,
	}
}
