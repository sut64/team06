package entity

import (
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
