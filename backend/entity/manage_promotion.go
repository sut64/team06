package entity

import (
	"time"
	"gorm.io/gorm"
)

type PromotionType struct {
	gorm.Model

	Type string

	//ManagePromotions []ManagePromotion `gorm:"foreignKey:PromotionTypeID"`
}
type NamePromotion struct {
	gorm.Model

	Name   string
	Detail string

	//ManagePromotions []ManagePromotion `gorm:"foreignKey:NamePromotionID"`
}
type PromotionPeriod struct {
	gorm.Model

	StartDate time.Time
	EndDate   time.Time

	//ManagePromotions []ManagePromotion `gorm:"foreignKey:PromotionPeriodID"`
}
