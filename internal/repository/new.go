package repository

import (
	"context"
	"gin-boilerplate/internal/repository/account"
	"gin-boilerplate/internal/repository/transactions"

	"gorm.io/gorm"
)

type Registry interface {
	Migrate() error

	DoInTx(ctx context.Context, fn func(ctx context.Context, tx Registry) error) error

	Account() account.IAccountRepo
	Transactions() transactions.ITransactionRepo
}

type impl struct {
	db *gorm.DB

	accountRepo account.IAccountRepo
	txRepo      transactions.ITransactionRepo
}

func New(dsn string, replicaDsn string) (Registry, error) {
	db, err := DbConnection(dsn, replicaDsn)
	if err != nil {
		return nil, err
	}

	return newImpl(db), nil
}

func newImpl(db *gorm.DB) Registry {
	return &impl{
		db:          db,
		accountRepo: account.New(db),
		txRepo:      transactions.New(db),
	}
}

func (i *impl) Account() account.IAccountRepo {
	return i.accountRepo
}

func (i *impl) Transactions() transactions.ITransactionRepo {
	return i.txRepo
}

func (i impl) DoInTx(ctx context.Context, fn func(ctx context.Context, tx Registry) error) error {
	tx := i.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(ctx, newImpl(tx)); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
