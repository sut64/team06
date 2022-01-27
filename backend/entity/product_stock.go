package entity

import (
	"gorm.io/gorm"
	"time"
)

type Productstock struct {
	gorm.Model
	Amount_remain   uint
	Update_datetime time.Time
	Detail          string 

	ProductID *uint
	Product   Product	`gorm:"foreignKey:ProductID"`

	ShelfstoreID *uint
	Shelfstore   Shelfstore	`gorm:"foreignKey:ShelfstoreID"`

	EmployeeID *uint
	Employee   Employee		`gorm:"foreignKey:EmployeeID"`
}

type Shelfstore struct {
	gorm.Model
	Zone string

	Productstock []Productstock	`gorm:"foreignKey:ProductstockID"`
}

type Product struct {
	gorm.Model
	Name  string
	Price float32

	Productstock []Productstock	`gorm:"foreignKey:ProductstockID"`

	TypeproductID *uint
	Typeproduct   Typeproduct	`gorm:"foreignKey:TypeproductID"`
}


type Typeproduct struct {
	gorm.Model
	Name string
	
	Product []Product	`gorm:"foreignKey:ProductID"`
}