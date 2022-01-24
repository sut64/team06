package entity

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	RoleName string

	UserLogins []UserLogin `gorm:"foreignKey:UserRoleID"`
}

type UserLogin struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string

	UserRoleID *uint    `gorm:"NOT NULL"`
	UserRole   UserRole `gorm:"references:ID; NOT NULL"`
}

type Gender struct {
	gorm.Model
	GenderName string

	UserDetails []UserDetail `gorm:"foreignKey:GenderID"`
}

type UserPrefix struct {
	gorm.Model
	PrefixName string

	UserDetails []UserDetail `gorm:"foreignKey:PrefixID"`
}

type UserDetail struct {
	gorm.Model

	PrefixID *uint
	Prefix   UserPrefix `gorm:"references:ID"`

	FirstName   string
	LastName    string
	PersonalID  string `gorm:"uniqueIndex"`
	PhoneNumber string
	Address     string

	GenderID *uint
	Gender   Gender `gorm:"references:ID"`
}
