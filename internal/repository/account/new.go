package account

import (
	"context"
	"gin-boilerplate/internal/models"

	"gorm.io/gorm"
)

type IAccountRepo interface {
	GetTxDetail(ctx context.Context, id int) (*models.Transaction, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) IAccountRepo {
	return &accountRepo{
		db: db,
	}
}
