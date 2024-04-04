package transactions

import (
	"context"
	"gin-boilerplate/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (i impl) CreateTransaction(ctx context.Context, s *models.Transaction) (*models.Transaction, error) {
	// Create transaction
	tx := models.Transaction{
		UUID:      uuid.New().String(),
		AccountID: s.AccountID,
		Amount:    s.Amount,
		TypeID:    s.TypeID,
		StatusID:  s.StatusID,
		CreatedAt: time.Now(),
	}

	if err := i.db.Create(&tx).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return &tx, nil
}
