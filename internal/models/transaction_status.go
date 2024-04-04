package models

import "gorm.io/gorm"

type TransactionStatus struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type TransactionStatusID uint

var (
	TransactionStatusCreated    TransactionStatusID = 1
	TransactionStatusProcessing TransactionStatusID = 2
	TransactionStatusSuccess    TransactionStatusID = 3
	TransactionStatusFailed     TransactionStatusID = 4
)
