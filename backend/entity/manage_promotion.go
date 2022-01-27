package entity

import (
	"time"

	"gorm.io/gorm"
)

type PromotionType struct {
	gorm.Model

	Type string

	ManagePromotions []ManagePromotion `gorm:"foreignKey:PromotionTypeID"`
}
type NamePromotion struct {
	gorm.Model

	Name   string
	Detail string

	ManagePromotions []ManagePromotion `gorm:"foreignKey:NamePromotionID"`
}
type PromotionPeriod struct {
	gorm.Model

	StartDate time.Time
	EndDate   time.Time

	ManagePromotions []ManagePromotion `gorm:"foreignKey:PromotionPeriodID"`
}
type ManagePromotion struct {
	gorm.Model

	EmployeeID *uint
	Employee   Employee `gorm:"references:ID"`

	PromotionPeriodID *uint
	PromotionPeriod   PromotionPeriod `gorm:"references:ID"`

	NamePromotionID *uint
	NamePromotion   NamePromotion `gorm:"references:ID"`

	PromotionTypeID *uint
	PromotionType   PromotionType `gorm:"references:ID"`

	PromotionCode string `gorm:"uniqueIndex"`
	MinPrice      float64

	Discount       float64
	Createdatetime time.Time
}
