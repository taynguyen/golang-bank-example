package models

import (
	"gorm.io/gorm"
)

type AccountStatus struct {
	gorm.Model
	Status string
}
