package transactions

import (
	"context"
	"gin-boilerplate/internal/models"

	"gorm.io/gorm"
)

type ITransactionRepo interface {
	GetTransactions(ctx context.Context, filter GetTransactionsFilter) ([]models.Transaction, models.Pagination, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
}

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITransactionRepo {
	return &impl{
		db: db,
	}
}
