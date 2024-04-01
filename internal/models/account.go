package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID          int64 `gorm:"index"`
	AccountStatusID int   `gorm:"index"`
	Number          string
	Balance         float64

	// Foreign keys
	AccountStatus AccountStatus
	User          User
}
