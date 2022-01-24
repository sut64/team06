package entity

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleName string

	// UserLogins []UserLogin `gorm:"foreignKey:UserRoleID"`
}

type UserLogin struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string

	UserRoleID *uint    `gorm:"NOT NULL"`
	UserRole   UserRole `gorm:"references:ID; NOT NULL"`
}
