package repository

import (
	"gin-boilerplate/internal/repository/account"
	"gin-boilerplate/internal/repository/transactions"

	"gorm.io/gorm"
)

type Registry interface {
	Migrate() error

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

	return &impl{
		db:          db,
		accountRepo: account.NewAccountRepo(db),
	}, nil
}

func (i *impl) Account() account.IAccountRepo {
	return i.accountRepo
}

func (i *impl) Transactions() transactions.ITransactionRepo {
	return i.txRepo
}
