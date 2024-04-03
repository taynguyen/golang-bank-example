package models

import "gorm.io/gorm"

type TransactionStatus struct {
	gorm.Model
	Name string `gorm:"not null"`
}
