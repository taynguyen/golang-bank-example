package accounts

import (
	"context"
	"gin-boilerplate/infra/logger"
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository"

	"gorm.io/gorm"
)

func (i impl) CreateTransaction(ctx context.Context, userID uint, t models.Transaction) (*models.Transaction, error) {
	logger := logger.GetLogger().WithField("controller", "CreateTransaction")
	t.StatusID = models.TransactionStatusCreated

	var newTx *models.Transaction
	// Create transaction
	if err := i.repo.DoInTx(ctx, func(ctx context.Context, tx repository.Registry) error {
		amount := float64(t.Amount)
		acc, err := tx.Account().GetAccount(ctx, t.AccountID)
		if err != nil {
			logger.Errorf("failed to get account: %v", err)
			return err
		}
		if acc == nil {
			logger.Errorf("account not found")
			return ErrAccountNotFound
		}

		// Check account balance again in transaction
		if t.TypeID == models.TransactionTypeWithdraw {
			if acc.Balance < amount {
				logger.Infof("insufficient balance")
				return ErrInsufficientBalance
			}

			// Withdraw
			amount = -amount
		}

		// Update balance
		if err := tx.Account().UpdateAccountBalance(ctx, &models.Account{
			Model: gorm.Model{
				ID: acc.ID,
			},
			Balance: acc.Balance + amount,
		}); err != nil {
			logger.Errorf("failed to update account balance: %v", err)
			return err
		}

		// Create transaction
		createdTx, err := tx.Transactions().CreateTransaction(ctx, &t)
		if err != nil {
			logger.Errorf("failed to create transaction: %v", err)
			return err
		}
		newTx = createdTx

		return nil

	}); err != nil {
		return nil, err
	}

	return newTx, nil
}
