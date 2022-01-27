package entity

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	MethodName string

	Orders []PurchaseOrder `gorm:"foreignKey:PaymentMethodID"`
}

type PurchaseOrder struct {
	gorm.Model

	MemberID *uint
	Member   Member `gorm:"references:ID"`

	PromotionID *uint
	Promotion   ManagePromotion `gorm:"references:ID"`

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod `gorm:"references:ID"`

	DeliveryAddress string
	OrderTime       time.Time
	OrderDiscount   float64
	OrderTotalPrice float64

	OrderItems []PurchaseOrderItem `gorm:"foreignKey:OrderID; constraint:OnDelete:CASCADE"`
}

type PurchaseOrderItem struct {
	gorm.Model

	OrderID *uint
	Order   PurchaseOrder `gorm:"references:ID"`

	ProductstockID *uint
	Productstock   Productstock

	OrderAmount uint
	ItemPrice   float64
}
