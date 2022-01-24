package entity

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleName string

	// UserLogins []UserLogin `gorm:"foreignKey:UserRoleID"`
}
