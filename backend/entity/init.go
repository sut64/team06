package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	//////////////////////////////////////
	//		 MANAGEPROMOTION		   //
	/////////////////////////////////////

	var PromotionPeriod0 = PromotionPeriod{
		StartDate: time.Now(),
		EndDate:   time.Now(),
	}
	db.Model(&PromotionPeriod{}).Create(&PromotionPeriod0)
	var PromotionPeriod1 = PromotionPeriod{
		StartDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local),
	}
	db.Model(&PromotionPeriod{}).Create(&PromotionPeriod1)
	var PromotionPeriod2 = PromotionPeriod{
		StartDate: time.Date(2022, 4, 10, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2022, 4, 20, 0, 0, 0, 0, time.Local),
	}
	db.Model(&PromotionPeriod{}).Create(&PromotionPeriod2)

	var PromotionPeriod3 = PromotionPeriod{
		StartDate: time.Date(2022, 10, 25, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2022, 11, 5, 0, 0, 0, 0, time.Local),
	}
	db.Model(&PromotionPeriod{}).Create(&PromotionPeriod3)

	var PromotionPeriod4 = PromotionPeriod{
		StartDate: time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local),
		EndDate:   time.Date(2022, 8, 1, 0, 0, 0, 0, time.Local),
	}
	db.Model(&PromotionPeriod{}).Create(&PromotionPeriod4)

	var NamePromotion0 = NamePromotion{
		Name:   "ไม่ใช้โปรโมชั่น",
		Detail: "ดีลสุดคุ้ม สินค้าใน Farm mart ซื้อขั้นต่ำตามกำหนดรับส่วนลดไปเลย ",
	}
	db.Model(&NamePromotion{}).Create(&NamePromotion0)
	var NamePromotion1 = NamePromotion{
		Name:   "FLASH DEAL",
		Detail: "ดีลสุดคุ้ม สินค้าใน Farm mart ซื้อขั้นต่ำตามกำหนดรับส่วนลดไปเลย ",
	}
	db.Model(&NamePromotion{}).Create(&NamePromotion1)
	var NamePromotion2 = NamePromotion{
		Name:   "BLACK FRIDAY",
		Detail: "ลดจัดหนัก สินค้าใน Farm mart ลดโหดรอบเดียวจบ! ",
	}
	db.Model(&NamePromotion{}).Create(&NamePromotion2)

	var NamePromotion3 = NamePromotion{
		Name:   "BUY MORE SAVE MORE",
		Detail: "ลดเห็นๆ ยิ่งซื้อมาก ยิ่งลดมาก เมื่อซื้อสินค้าใน Farm mart  ",
	}
	db.Model(&NamePromotion{}).Create(&NamePromotion3)

	var PromotionType1 = PromotionType{
		Type: "ช่วงเทศกาล",
	}
	db.Model(&PromotionType{}).Create(&PromotionType1)

	var PromotionType2 = PromotionType{
		Type: "ตามราคาขั้นต่ำในบิล",
	}
	db.Model(&PromotionType{}).Create(&PromotionType2)

	ManagePromotion0 := ManagePromotion{
		PromotionCode:  "A0000",
		Employee:       employeeSakeet,
		NamePromotion:  NamePromotion0,
		MinPrice:       0,
		Discount:       0,
		Createdatetime: time.Now(),
	}
	db.Model(&ManagePromotion{}).Create(&ManagePromotion0)
	ManagePromotion1 := ManagePromotion{
		Employee:        employeeSakeet,
		PromotionCode:   "A1500",
		NamePromotion:   NamePromotion1,
		PromotionPeriod: PromotionPeriod1,
		PromotionType:   PromotionType1,
		MinPrice:        100,
		Discount:        10,
		Createdatetime:  time.Now(),
	}
	db.Model(&ManagePromotion{}).Create(&ManagePromotion1)

	/////////////////////////////////////
	//		    Productstock		   //
	/////////////////////////////////////

	typeproduct1 := Typeproduct{
		Name: "Snack",
	}
	db.Model(&Typeproduct{}).Create(&typeproduct1)

	typeproduct2 := Typeproduct{
		Name: "Food",
	}
	db.Model(&Typeproduct{}).Create(&typeproduct2)

	product1 := Product{
		Name:        "Lays",
		Price:       20.50,
		Typeproduct: typeproduct1,
	}
	db.Model(&Product{}).Create(&product1)

	shelfstore1 := Shelfstore{
		Zone: "A",
	}
	db.Model(&Shelfstore{}).Create(&shelfstore1)

	shelfstore2 := Shelfstore{
		Zone: "B",
	}
	db.Model(&Shelfstore{}).Create(&shelfstore2)

	productstock1 := Productstock{
		Amount_remain:   20,
		Update_datetime: time.Now(),
		Detail:          "snack for kid",
		Product:         product1,
		Employee:        employeeSakeet,
		Shelfstore:      shelfstore1,
	}
	db.Model(&Productstock{}).Create(&productstock1)

	/////////////////////////////////////
	//		    PurchaseOrder		   //
	/////////////////////////////////////
	cashMethod := PaymentMethod{
		MethodName: "เงินสด",
	}
	mobileMethod := PaymentMethod{
		MethodName: "Mobile banking",
	}
	bankMethod := PaymentMethod{
		MethodName: "หักผ่านบัญชี",
	}
	db.Model(&PaymentMethod{}).Create(&cashMethod)
	db.Model(&PaymentMethod{}).Create(&mobileMethod)
	db.Model(&PaymentMethod{}).Create(&bankMethod)

}
