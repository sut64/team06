package entity

import (
	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	DayOfWeek string

	// ManageWorkTime []ManageWorkTime `gorm:"foreignKey:DayID"`
}
