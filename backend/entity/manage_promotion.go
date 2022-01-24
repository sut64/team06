package entity

import (
	"time"

	"gorm.io/gorm"
)

type PromotionType struct {
	gorm.Model

	Type string

	WManagePromotions []ManagePromotion `gorm:"foreignKey:PromotionTypeID"`
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

	PromotionPeriodID *uint
	PromotionPeriod   PromotionPeriod `gorm:"references:ID"`

	NamePromotionID *uint
	NamePromotion   NamePromotion `gorm:"references:ID"`

	PromotionTypeID *uint
	PromotionType   PromotionType `gorm:"references:ID"`

	PromotionCode  string
	MinPrice       float64
	Quantity       uint
	Discount       float64
	Createdatetime time.Time
}
