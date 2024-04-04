package account

import (
	"context"
	"gin-boilerplate/internal/models"

	"gorm.io/gorm"
)

type IAccountRepo interface {
	GetAccount(ctx context.Context, id uint) (*models.Account, error)
	UpdateAccountBalance(ctx context.Context, account *models.Account) error

	GetTxDetail(ctx context.Context, id int) (*models.Transaction, error)
}

type impl struct {
	db *gorm.DB
}

func New(db *gorm.DB) IAccountRepo {
	return &impl{
		db: db,
	}
}
