package repository

import (
	"gin-boilerplate/internal/repository/account"

	"gorm.io/gorm"
)

type Registry interface {
	Migrate() error

	Account() account.IAccountRepo
}

type impl struct {
	db *gorm.DB

	account account.IAccountRepo
}

func New(dsn string, replicaDsn string) (Registry, error) {
	db, err := DbConnection(dsn, replicaDsn)
	if err != nil {
		return nil, err
	}

	return &impl{
		db:      db,
		account: account.NewAccountRepo(db),
	}, nil
}

func (i *impl) Account() account.IAccountRepo {
	return i.account
}
