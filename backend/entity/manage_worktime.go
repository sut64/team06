package entity

import (
	"time"

	"gorm.io/gorm"
)

type ManageWorkTime struct {
	gorm.Model
	Comment     string
	WorkingDate time.Time
	TimeTotal   uint

	ManagerID *uint
	Manager   Employee

	EmployeeID *uint
	Employee   Employee

	DayID *uint
	Day   Day

	MonthID *uint
	Month   Month

	WorkingTimeID *uint
	WorkingTime   WorkingTime

	ManageSalary []ManageSalary `gorm:"foreignkey:ManageWorkTimeID"`
}

type Day struct {
	gorm.Model
	DayNumber uint

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:DayID"`
}

type Month struct {
	gorm.Model
	MonthOfYear string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:MonthID"`
}

type WorkingTime struct {
	gorm.Model
	TimeToTime string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WorkingTimeID"`
}
