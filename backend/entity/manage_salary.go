package entity

import (
	"time"

	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	Level uint
	Name  string
}

type BonusStatus struct {
	gorm.Model
	Name string
}
