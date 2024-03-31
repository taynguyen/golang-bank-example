package repository

import (
	"gin-boilerplate/internal/models"
)

func (i impl) Migrate() error {
	var migrationModels = []interface{}{&models.Example{}}
	return i.db.AutoMigrate(migrationModels...)
}
