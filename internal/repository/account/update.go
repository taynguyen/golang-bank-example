package account

import (
	"context"
	"gin-boilerplate/internal/models"
)

func (i impl) UpdateAccountBalance(ctx context.Context, account *models.Account) error {
	return i.db.Model(&models.Account{}).
		Where("id = ?", account.ID).
		Update("balance", account.Balance).
		Update("updated_at", account.UpdatedAt).Error
}
