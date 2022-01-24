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
	Product   Product

	ShelfstoreID *uint
	Shelfstore   Shelfstore

	EmployeeID *uint
	Employee   Employee
}


type Typeproduct struct {
	gorm.Model
	Name string
	
	Product []Product
}