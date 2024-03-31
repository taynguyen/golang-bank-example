package migrations

import (
	"gin-boilerplate/internal/models"
	"gin-boilerplate/internal/repository"
)

// Migrate Add list of model add for migrations
// TODO later separate migration each models
func Migrate() error {
	var migrationModels = []interface{}{&models.Example{}}
	return repository.DB.AutoMigrate(migrationModels...)
}
