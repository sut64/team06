package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

var managerMai = Employee{
	UserLogin: UserLogin{
		Username: "M6225605",
		Password: SetupPasswordHash("M6225605"),
		UserRole: UserRole{
			RoleName:   "Employee",
			RoleNameTH: "พนักงาน",
		},
	},
	UserDetail: UserDetail{
		Prefix: UserPrefix{
			PrefixName: "นาย",
		},
		FirstName:   "ณัฐวัตร",
		LastName:    "บุญโสดากร",
		PersonalID:  "5236587453655",
		PhoneNumber: "0615852336",
		Address:     "หอพักสุรนิเวศ 7, มทส.",
		Gender: Gender{
			GenderName: "ชาย",
		},
	},
	Position: EmployeePosition{
		PositionName:   "Manager",
		PositionNameTH: "ผู้จัดการ",
		Salary:         50000.0,
	},
}

var ManageWT = ManageWorkTime{
	NameSchedule: "ผู้จัดการ",
	WorkingDate:  time.Date(2021, 1, 1, 9, 0, 0, 0, time.UTC),
	TimeTotal:    9,
	Manager:      managerMai,
	Employee:     managerMai,
	Day: Day{
		DayOfWeek: "วันศุกร์",
	},
	Weekly: Weekly{
		WeekAt: 1,
	},
	Month: Month{
		MonthOfYear: "มกราคม",
	},
	WorkingTime: WorkingTime{
		TimeToTime: "08.00-17.00",
	},
}

var Assessment_5 = Assessment{
	Level: 5,
	Name:  "ดีเยี่ยม",
}

var BonusStatus_approve = BonusStatus{
	Name: "อนุมัติ",
}

func TestManageSalaryPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	manageSalary := ManageSalary{
		Manager:        managerMai,
		ManageWorkTime: ManageWT,
		Assessment:     Assessment_5,
		BonusAmount:    2000.00,
		BonusDetail:    "ทำงานรวดเร็ว",
		BonusStatus:    BonusStatus_approve,
		CreateAt:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageSalary)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestManageSalaryBonusAmountMustBePositive(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []float64{
		-100,
		-200.00,
	}

	// ข้อมูลถูกต้องหมดทุก field
	for _, fixture := range fixtures {
		manageSalary := ManageSalary{
			Manager:        managerMai,
			ManageWorkTime: ManageWT,
			Assessment:     Assessment_5,
			BonusAmount:    fixture, // ผิด
			BonusDetail:    "ทำงานรวดเร็ว",
			BonusStatus:    BonusStatus_approve,
			CreateAt:       time.Now(),
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(manageSalary)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("BonusAmount must be positive number"))
	}
}

func TestManageSalaryBonusDetailMustBeValid(t *testing.T) {
	g := NewGomegaWithT(t)

	test_data := ""
	for i := 0; i < 201; i++ {
		test_data += "A"
	}

	// ข้อมูลถูกต้องหมดทุก field
	manageSalary := ManageSalary{
		Manager:        managerMai,
		ManageWorkTime: ManageWT,
		Assessment:     Assessment_5,
		BonusAmount:    2000.00,
		BonusDetail:    test_data, // ผิด
		BonusStatus:    BonusStatus_approve,
		CreateAt:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageSalary)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("BonusDetail length must be between 0 - 200"))
}

func TestManageSalaryCreateAtMustBePresent(t *testing.T) {
	g := NewGomegaWithT(t)

	test_data := time.Now().Add(-24 * time.Hour)

	// ข้อมูลถูกต้องหมดทุก field
	manageSalary := ManageSalary{
		Manager:        managerMai,
		ManageWorkTime: ManageWT,
		Assessment:     Assessment_5,
		BonusAmount:    2000.00,
		BonusDetail:    "ทำงานรวดเร็ว",
		BonusStatus:    BonusStatus_approve,
		CreateAt:       test_data, // ผิด
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageSalary)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("CreateAt must be present"))
}
