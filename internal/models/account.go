package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID          uint `gorm:"index"`
	AccountStatusID int  `gorm:"index"`
	BankID          int
	Number          string
	Balance         float64

	// Foreign keys
	AccountStatus *AccountStatus
	Bank          *Bank
	User          *User
}
