package accounts

import (
	"context"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository"
	"gin-boilerplate/internal/repository/transactions"
)

type IAccountController interface {
	GetTransactions(ctx context.Context, filter transactions.GetTransactionsFilter) ([]models.Transaction, models.Pagination, error)
	CreateTransaction(ctx context.Context, userID uint, transaction models.Transaction) (*models.Transaction, error)
}

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) IAccountController {
	return &impl{
		repo: repo,
	}
}
