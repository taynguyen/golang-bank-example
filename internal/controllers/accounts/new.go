package accounts

import (
	"context"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository"
)

type IAccountController interface {
	GetTransactions(ctx context.Context, userID int64) ([]models.Transaction, models.Pagination, error)
}

type impl struct {
	repo repository.Registry
}

func New(repo repository.Registry) IAccountController {
	return &impl{
		repo: repo,
	}
}
