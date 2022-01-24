package entity

import (
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

func SetupPasswordHash(pwd string) string {
	var password, _ = bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(password)
}

func SetupIntoDatabase(db *gorm.DB) {
	/////////////////////////////////////
	//				USER			   //
	/////////////////////////////////////
	// UserRole
	customerRole := UserRole{
		RoleName:   "Member",
		RoleNameTH: "สมาชิก",
	}
	employeeRole := UserRole{
		RoleName:   "Employee",
		RoleNameTH: "พนักงาน",
	}
	db.Model(&UserRole{}).Create(&customerRole)
	db.Model(&UserRole{}).Create(&employeeRole)

	// Gender
	maleGender := Gender{
		GenderName: "ชาย",
	}
	femaleGender := Gender{
		GenderName: "หญิง",
	}
	otherGender := Gender{
		GenderName: "อื่น ๆ",
	}
	notSpecifiedGender := Gender{
		GenderName: "ไม่ระบุ",
	}
	db.Model(&Gender{}).Create(&maleGender)
	db.Model(&Gender{}).Create(&femaleGender)
	db.Model(&Gender{}).Create(&otherGender)
	db.Model(&Gender{}).Create(&notSpecifiedGender)

	// Prefix
	boyPrefix := UserPrefix{
		PrefixName: "เด็กชาย",
	}
	girlPrefix := UserPrefix{
		PrefixName: "เด็กหญิง",
	}
	mrPrefix := UserPrefix{
		PrefixName: "นาย",
	}
	mrsPrefix := UserPrefix{
		PrefixName: "นาง",
	}
	missPrefix := UserPrefix{
		PrefixName: "นางสาว",
	}
	db.Model(&UserPrefix{}).Create(&boyPrefix)
	db.Model(&UserPrefix{}).Create(&girlPrefix)
	db.Model(&UserPrefix{}).Create(&mrPrefix)
	db.Model(&UserPrefix{}).Create(&mrsPrefix)
	db.Model(&UserPrefix{}).Create(&missPrefix)

	// UserLogin
	/* **** Can edit this later **** */
	login1 := UserLogin{
		Username: "foei",
		Password: SetupPasswordHash("1234"),
		UserRole: customerRole,
	}
	login2 := UserLogin{
		Username: "sakeet",
		Password: SetupPasswordHash("5678"),
		UserRole: employeeRole,
	}
	login3 := UserLogin{
		Username: "cue",
		Password: SetupPasswordHash("2468"),
		UserRole: employeeRole,
	}
	db.Model(&UserLogin{}).Create(&login1)
	db.Model(&UserLogin{}).Create(&login2)
	db.Model(&UserLogin{}).Create(&login3)

	// UserDetail
	userDetail1 := UserDetail{
		Prefix:      mrPrefix,
		FirstName:   "วณัฐพงศ์",
		LastName:    "วงษ์บู่ทอง",
		PersonalID:  "141xx018xx5xx",
		PhoneNumber: "0862275814",
		Address:     "หอพักสุรนิเวศ 9, มทส.",
		Gender:      maleGender,
	}
	userDetail2 := UserDetail{
		Prefix:      mrPrefix,
		FirstName:   "สหรัฐ",
		LastName:    "จันทรินทร์",
		PersonalID:  "1234567890123",
		PhoneNumber: "0921548756",
		Address:     "บ้านปณิธาน, นครราชสีมา",
		Gender:      maleGender,
	}
	userDetail3 := UserDetail{
		Prefix:      mrPrefix,
		FirstName:   "ฝนเทพ",
		LastName:    "พลวัต",
		PersonalID:  "4567890123456",
		PhoneNumber: "0847852563",
		Address:     "หมู่บ้านสุรสวัสดิ์, นครราชสีมา",
		Gender:      maleGender,
	}
	db.Model(&UserDetail{}).Create(&userDetail1)
	db.Model(&UserDetail{}).Create(&userDetail2)
	db.Model(&UserDetail{}).Create(&userDetail3)

	// EmployeePosition
	employeePos := EmployeePosition{
		PositionName:   "Employee",
		PositionNameTH: "พนักงาน",
		Salary:         25000.0,
	}
	managerPos := EmployeePosition{
		PositionName:   "Manager",
		PositionNameTH: "ผู้จัดการ",
		Salary:         50000.0,
	}
	db.Model(&EmployeePosition{}).Create(&employeePos)
	db.Model(&EmployeePosition{}).Create(&managerPos)

	// Member
	memberFoei := Member{
		UserLogin:  login1,
		UserDetail: userDetail1,
	}
	db.Model(&Member{}).Create(&memberFoei)
	// Manager
	managerCue := Employee{
		UserLogin:  login3,
		UserDetail: userDetail3,
		Position:   managerPos,
	}
	db.Model(&Employee{}).Create(&managerCue)
	// Employee
	employeeSakeet := Employee{
		UserLogin:  login2,
		UserDetail: userDetail2,
		Position:   employeePos,
		Superior:   &managerCue,
	}
	db.Model(&Employee{}).Create(&employeeSakeet)
}
