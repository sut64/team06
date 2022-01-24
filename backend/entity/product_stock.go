package entity

import (
	"gorm.io/gorm"
)


type Typeproduct struct {
	gorm.Model
	Name string
}