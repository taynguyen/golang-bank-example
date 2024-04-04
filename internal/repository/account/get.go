package account

import (
	"context"
	"gin-boilerplate/internal/models"
)

func (i *impl) GetTxDetail(ctx context.Context, id int) (*models.Transaction, error) {
	var tx models.Transaction
	err := i.db.WithContext(ctx).Where("id = ?", id).First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (i *impl) GetAccount(ctx context.Context, id uint) (*models.Account, error) {
	var account models.Account
	err := i.db.WithContext(ctx).Where("id = ?", id).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}
