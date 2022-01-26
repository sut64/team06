package entity

import (
	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	DayOfWeek string

	// ManageWorkTime []ManageWorkTime `gorm:"foreignKey:DayID"`
}

type Month struct {
	gorm.Model
	MonthOfYear string

	// ManageWorkTime []ManageWorkTime `gorm:"foreignKey:MonthID"`
}

type Weekly struct {
	gorm.Model
	WeekAt uint

	// ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WeeklyID"`
}

type WorkingTime struct {
	gorm.Model
	TimeToTime string

	// ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WorkingTimeID"`
}
