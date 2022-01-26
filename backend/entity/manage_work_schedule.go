package entity

import (
	"time"

	"gorm.io/gorm"
)

type ManageWorkTime struct {
	gorm.Model
	NameSchedule string
	WorkingDate  time.Time
	TimeTotal    uint

	ManagerID *uint
	Manager   Manager

	EmployeeID *uint
	Employee   Employee

	DayID *uint
	Day   Day

	WeeklyID *uint
	Weekly   Weekly

	MonthID *uint
	Month   Month

	WorkingTimeID *uint
	WorkingTime   WorkingTime

	// ManageSalary []ManageSalary `gorm:"foreignkey:ManageWorkTimeID"`
}

type Day struct {
	gorm.Model
	DayOfWeek string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:DayID"`
}

type Month struct {
	gorm.Model
	MonthOfYear string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:MonthID"`
}

type Weekly struct {
	gorm.Model
	WeekAt uint

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WeeklyID"`
}

type WorkingTime struct {
	gorm.Model
	TimeToTime string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WorkingTimeID"`
}

type Manager struct {
	gorm.Model
	Name     string
	Code     string
	Password string
	Employee []Employee `gorm:"foreignKey:ManagerID"`
}
