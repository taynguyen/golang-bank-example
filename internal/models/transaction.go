package models

import "time"

type Transaction struct {
	UUID      string  `gorm:"primaryKey"`
	AccountID uint    `gorm:"not null"`
	TypeID    uint    `gorm:"not null"`
	Amount    float32 `gorm:"not null"`
	StatusID  uint    `gorm:"not null"`

	CreatedAt time.Time `gorm:"index:,sort:desc"`
	UpdatedAt time.Time

	// Foreign keys
	Type   *TransactionType
	Status *TransactionStatus
}
