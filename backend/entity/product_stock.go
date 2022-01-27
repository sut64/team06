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

	Product   Product	`gorm:"references:ID"`

	ShelfstoreID *uint
	Shelfstore   Shelfstore	`gorm:"references:ID"`

	EmployeeID *uint
	Employee   Employee		`gorm:"references:ID"`

}

type Shelfstore struct {
	gorm.Model
	Zone string


	Productstock []Productstock	`gorm:"foreignkey:ShelfstoreID"`

}

type Product struct {
	gorm.Model
	Name  string
	Price float32


	Productstock []Productstock	`gorm:"foreignkey:ProductID"`

	TypeproductID *uint
	Typeproduct   Typeproduct	`gorm:"references:ID"`

}

type Typeproduct struct {
	gorm.Model
	Name string

	
	Product []Product	`gorm:"foreignkey:TypeproductID"`
}

