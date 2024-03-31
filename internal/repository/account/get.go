package account

import (
	"context"
	"gin-boilerplate/internal/models"
)

func (i *accountRepo) GetTxDetail(ctx context.Context, id int) (*models.Transaction, error) {
	var tx models.Transaction
	err := i.db.Where("id = ?", id).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}
