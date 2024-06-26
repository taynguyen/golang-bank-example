package repository

import "gin-boilerplate/internal/models"

func (i impl) Migrate() error {
	var migrationModels = []interface{}{
		&models.Bank{},
		&models.User{},
		&models.Account{},
		&models.AccountStatus{},
		&models.Transaction{},
		&models.TransactionType{},
	}

	return i.db.AutoMigrate(migrationModels...)
}
