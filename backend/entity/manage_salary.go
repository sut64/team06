package entity

import (
	"time"

	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	Level uint
	Name  string

	ManageSalarys []ManageSalary `gorm:"foreignkey:AssessmentID"`
}

type BonusStatus struct {
	gorm.Model
	Name string

	ManageSalarys []ManageSalary `gorm:"foreignkey:BonusStatusID"`
}

type ManageSalary struct {
	gorm.Model

	ManagerID *uint
	Manager   Employee `gorm:"references:ID"`

	ManageWorkTimeID *uint
	ManageWorkTime   ManageWorkTime `gorm:"references:ID"`

	AssessmentID *uint
	Assessment   Assessment `gorm:"references:ID"`

	BonusAmount float64
	BonusDetail string

	BonusStatusID *uint
	BonusStatus   BonusStatus `gorm:"references:ID"`

	CreateAt time.Time
}
