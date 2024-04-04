package models

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type TransactionTypeID = uint

const (
	TransactionTypeWithdraw = TransactionTypeID(1)
	TransactionTypeDeposit  = TransactionTypeID(2)
)
