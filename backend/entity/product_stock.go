package entity

import (
	"time"

	"gorm.io/gorm"
)

type Productstock struct {
	gorm.Model
	Amount_remain   uint
	Update_datetime time.Time
	Detail          string

	ProductID *uint
	Product   Product

	ShelfstoreID *uint
	Shelfstore   Shelfstore

	EmployeeID *uint
	Employee   Employee

	OrderItems []PurchaseOrderItem `gorm:"foreignKey:ProductstockID"`
}

type Shelfstore struct {
	gorm.Model
	Zone string

	Productstock []Productstock
}

type Product struct {
	gorm.Model
	Name  string
	Price float32

	Productstock []Productstock

	TypeproductID *uint
	Typeproduct   Typeproduct
}

type Typeproduct struct {
	gorm.Model
	Name string

	Product []Product
}
