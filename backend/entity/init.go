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
		Username: "C6220709",
		Password: SetupPasswordHash("C6220709"),
		UserRole: customerRole,
	}
	login2 := UserLogin{
		Username: "E6209568",
		Password: SetupPasswordHash("E6209568"),
		UserRole: employeeRole,
	}
	login3 := UserLogin{
		Username: "E6230258",
		Password: SetupPasswordHash("E6230258"),
		UserRole: employeeRole,
	}
	login4 := UserLogin{
		Username: "M6225605",
		Password: SetupPasswordHash("M6225605"),
		UserRole: employeeRole,
	}
	login5 := UserLogin{
		Username: "C6111618",
		Password: SetupPasswordHash("C6111618"),
		UserRole: customerRole,
	}
	login6 := UserLogin{
		Username: "M6026400",
		Password: SetupPasswordHash("M6026400"),
		UserRole: employeeRole,
	}
	db.Model(&UserLogin{}).Create(&login1)
	db.Model(&UserLogin{}).Create(&login2)
	db.Model(&UserLogin{}).Create(&login3)
	db.Model(&UserLogin{}).Create(&login4)
	db.Model(&UserLogin{}).Create(&login5)
	db.Model(&UserLogin{}).Create(&login6)

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
	userDetail4 := UserDetail{
		Prefix:      mrPrefix,
		FirstName:   "ณัฐวัตร",
		LastName:    "บุญโสดากร",
		PersonalID:  "5236587453655",
		PhoneNumber: "0615852336",
		Address:     "หอพักสุรนิเวศ 7, มทส.",
		Gender:      maleGender,
	}
	userDetail5 := UserDetail{
		Prefix:      mrPrefix,
		FirstName:   "เศรษฐกฤษ",
		LastName:    "โพธิ์แก้ว",
		PersonalID:  "7885631455574",
		PhoneNumber: "0842221960",
		Address:     "กรุงเทพมหานคร",
		Gender:      maleGender,
	}
	userDetail6 := UserDetail{
		Prefix:      missPrefix,
		FirstName:   "ชัฎฌาฎา",
		LastName:    "เรือนทอง",
		PersonalID:  "4110075486354",
		PhoneNumber: "0811917200",
		Address:     "หอพักสุรนิเวศ 15, มทส.",
		Gender:      femaleGender,
	}
	db.Model(&UserDetail{}).Create(&userDetail1)
	db.Model(&UserDetail{}).Create(&userDetail2)
	db.Model(&UserDetail{}).Create(&userDetail3)
	db.Model(&UserDetail{}).Create(&userDetail4)
	db.Model(&UserDetail{}).Create(&userDetail5)
	db.Model(&UserDetail{}).Create(&userDetail6)

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
	memberKobkab := Member{
		UserLogin:  login5,
		UserDetail: userDetail5,
	}
	db.Model(&Member{}).Create(&memberFoei)
	db.Model(&Member{}).Create(&memberKobkab)
	// Manager
	managerMai := Employee{
		UserLogin:  login4,
		UserDetail: userDetail4,
		Position:   managerPos,
	}
	managerEarn := Employee{
		UserLogin:  login6,
		UserDetail: userDetail6,
		Position:   managerPos,
	}
	db.Model(&Employee{}).Create(&managerMai)
	db.Model(&Employee{}).Create(&managerEarn)
	// Employee
	employeeSakeet := Employee{
		UserLogin:  login2,
		UserDetail: userDetail2,
		Position:   employeePos,
		Superior:   &managerMai,
	}
	employeeCue := Employee{
		UserLogin:  login3,
		UserDetail: userDetail3,
		Position:   employeePos,
		Superior:   &managerEarn,
	}
	db.Model(&Employee{}).Create(&employeeSakeet)
	db.Model(&Employee{}).Create(&employeeCue)

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
		Name:        "Testo 50g",
		Price:       20,
		Typeproduct: typeproduct1,
	}
	db.Model(&Product{}).Create(&product1)

	product2 := Product{
		Name:        "Coconut 1pc",
		Price:       29,
		Typeproduct: typeproduct2,
	}
	db.Model(&Product{}).Create(&product2)

	product3 := Product{
		Name:        "Fresh Medium No. 3-4 Fresh Egg 30pcs",
		Price:       110,
		Typeproduct: typeproduct2,
	}
	db.Model(&Product{}).Create(&product3)

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

	/////////////////////////////////////
	//		    ManageWorkTime		   //
	/////////////////////////////////////
	// Add day
	allDay := []string{"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"}
	for _, s := range allDay {
		tDay := Day{
			DayOfWeek: s,
		}
		db.Model(&Day{}).Create(&tDay)
	}

	// Add weekly
	allWeek := []uint{1, 2, 3, 4}
	for _, i := range allWeek {
		tWeek := Weekly{
			WeekAt: i,
		}
		db.Model(&Weekly{}).Create(&tWeek)
	}

	// Add month
	allWMonth := []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษถาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"}
	for _, m := range allWMonth {
		tMonth := Month{
			MonthOfYear: m,
		}
		db.Model(&Month{}).Create(&tMonth)
	}

	// Add working time
	allWorkingTime := []string{"08.00-17.00", "09.00-18.00", "10.00-19.00", "11.00-20.00", "12.00-21.00", "13.00-22.00", "14.00-23.00"}
	for _, t := range allWorkingTime {
		tWork := WorkingTime{
			TimeToTime: t,
		}
		db.Model(&WorkingTime{}).Create(&tWork)
	}

	/////////////////////////////////////
	//		    ManageSalary		   //
	/////////////////////////////////////
	// Assessment data
	Assessment_1 := Assessment{
		Level: 1,
		Name:  "แย่",
	}
	Assessment_2 := Assessment{
		Level: 2,
		Name:  "ปรับปรุง",
	}
	Assessment_3 := Assessment{
		Level: 3,
		Name:  "พอใช้",
	}
	Assessment_4 := Assessment{
		Level: 4,
		Name:  "ดี",
	}
	Assessment_5 := Assessment{
		Level: 5,
		Name:  "ดีเยี่ยม",
	}

	// BonusStatus data
	BonusStatus_wait := BonusStatus{
		Name: "รออนุมัติ",
	}
	BonusStatus_approve := BonusStatus{
		Name: "อนุมัติ",
	}
	BonusStatus_disapproved := BonusStatus{
		Name: "ไม่อนุมัติ",
	}

	// ManageWorkTime data
	var Day_Fr Day
	db.Model(&Day{}).Find(&Day_Fr, db.Where("id = ?", 6))

	var Week_01 Weekly
	db.Model(&Weekly{}).Find(&Week_01, db.Where("id = ?", 1))

	var Month_1Jan Month
	db.Model(&Month{}).Find(&Month_1Jan, db.Where("id = ?", 1))

	var wt_8t17 WorkingTime
	db.Model(&WorkingTime{}).Find(&wt_8t17, db.Where("id = ?", 1))

	ManageWT_M01 := ManageWorkTime{
		NameSchedule:  "ผู้จัดการ",
		WorkingDate:   time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC),
		TimeTotal:     9,
		Manager:       managerMai,
		Employee:      managerMai,
		DayID:         &Day_Fr.ID,
		WeeklyID:      &Week_01.ID,
		MonthID:       &Month_1Jan.ID,
		WorkingTimeID: &wt_8t17.ID,
	}
	ManageWT_E01 := ManageWorkTime{
		NameSchedule:  "พนักงาน",
		WorkingDate:   time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC),
		TimeTotal:     8,
		Manager:       managerMai,
		Employee:      employeeSakeet,
		DayID:         &Day_Fr.ID,
		WeeklyID:      &Week_01.ID,
		MonthID:       &Month_1Jan.ID,
		WorkingTimeID: &wt_8t17.ID,
	}

	// ManageSalary data
	MS_01 := ManageSalary{
		Manager:        managerMai,
		ManageWorkTime: ManageWT_M01,
		AssessmentID:   &Assessment_5.ID,
		BonusAmount:    2000.00,
		BonusDetail:    "ทำงานรวดเร็ว",
		BonusStatusID:  &BonusStatus_approve.ID,
		CreateAt:       time.Date(2021, 1, 28, 10, 0, 0, 0, time.UTC),
	}
	MS_02 := ManageSalary{
		Manager:        managerMai,
		ManageWorkTime: ManageWT_E01,
		AssessmentID:   &Assessment_5.ID,
		BonusAmount:    1500.00,
		BonusDetail:    "ทำงานดี",
		BonusStatusID:  &BonusStatus_approve.ID,
		CreateAt:       time.Date(2021, 1, 28, 10, 0, 0, 0, time.UTC),
	}

	// add data
	db.Model(&Assessment{}).Create(&Assessment_1)
	db.Model(&Assessment{}).Create(&Assessment_2)
	db.Model(&Assessment{}).Create(&Assessment_3)
	db.Model(&Assessment{}).Create(&Assessment_4)
	db.Model(&Assessment{}).Create(&Assessment_5)
	db.Model(&BonusStatus{}).Create(&BonusStatus_wait)
	db.Model(&BonusStatus{}).Create(&BonusStatus_approve)
	db.Model(&BonusStatus{}).Create(&BonusStatus_disapproved)
	db.Model(&ManageSalary{}).Create(&MS_01)
	db.Model(&ManageSalary{}).Create(&MS_02)

}
