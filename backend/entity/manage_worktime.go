package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ManageWorkTime struct {
	gorm.Model
	Comment     string
	WorkingDate time.Time
	TimeTotal   uint `valid:"customEqualEight~TimeTotal must be equal 8 hr."`

	ManagerID *uint
	Manager   Employee `gorm:"references:ID" valid:"-"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:ID" valid:"-"`

	DayID *uint
	Day   Day `gorm:"references:ID" valid:"-"`

	MonthID *uint
	Month   Month `gorm:"references:ID" valid:"-"`

	WorkingTimeID *uint
	WorkingTime   WorkingTime `gorm:"references:ID" valid:"-"`

	ManageSalary []ManageSalary `gorm:"foreignkey:ManageWorkTimeID"`
}
type Day struct {
	gorm.Model
	DayNumber uint

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:DayID"`
}

type Month struct {
	gorm.Model
	MonthOfYear string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:MonthID"`
}

type WorkingTime struct {
	gorm.Model
	TimeToTime string

	ManageWorkTime []ManageWorkTime `gorm:"foreignKey:WorkingTimeID"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("customEqualEight", func(i interface{}, context interface{}) bool {
		return i.(uint) == 8
	})

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})
}
