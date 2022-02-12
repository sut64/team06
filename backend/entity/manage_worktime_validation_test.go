package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

var managerChut = Employee{
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
			PrefixName: "นางสาว",
		},
		FirstName:   "ชัฎฌาฎา",
		LastName:    "เรือนทอง",
		PersonalID:  "4110075486354",
		PhoneNumber: "0811917200",
		Address:     "หอพักสุรนิเวศ 15, มทส.",
		Gender: Gender{
			GenderName: "หญิง",
		},
	},
	Position: EmployeePosition{
		PositionName:   "Manager",
		PositionNameTH: "ผู้จัดการ",
		Salary:         50000.0,
	},
}

func TestManageWorkTimePass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	manageWorkTime := ManageWorkTime{
		Comment:     "ตารางงานพนักงาน",
		WorkingDate: time.Now(),
		TimeTotal:   8,
		Manager:     managerChut,
		Employee:    managerChut,
		Day: Day{
			DayNumber: 26,
		},
		Month: Month{
			MonthOfYear: "มกราคม",
		},
		WorkingTime: WorkingTime{
			TimeToTime: "08.00-17.00",
		},
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageWorkTime)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("TimeTotal must be equal 8 hr."))
}

func TestManageWorkTimeCommentMustBeValid(t *testing.T) {
	g := NewGomegaWithT(t)

	test_data := ""
	for i := 0; i < 201; i++ {
		test_data += "B"
	}

	// ข้อมูลถูกต้องหมดทุก field
	manageWorkTime := ManageWorkTime{
		Comment:     test_data,
		WorkingDate: time.Now(),
		TimeTotal:   8,
		Manager:     managerChut,
		Employee:    managerChut,
		Day: Day{
			DayNumber: 26,
		},
		Month: Month{
			MonthOfYear: "มกราคม",
		},
		WorkingTime: WorkingTime{
			TimeToTime: "08.00-17.00",
		},
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageWorkTime)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Comment length must be between 0 - 200"))
}

func TestManageWorkTimeTimeTotalMustEqualEight(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []uint{
		6,
		9,
	}

	// ข้อมูลถูกต้องหมดทุก field
	for _, fixture := range fixtures {
		manageWorkTime := ManageWorkTime{
			Comment:     "ตารางงานพนักงาน",
			WorkingDate: time.Now(),
			TimeTotal:   fixture, // ผิด
			Manager:     managerChut,
			Employee:    managerChut,
			Day: Day{
				DayNumber: 26,
			},
			Month: Month{
				MonthOfYear: "มกราคม",
			},
			WorkingTime: WorkingTime{
				TimeToTime: "08.00-17.00",
			},
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(manageWorkTime)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("TimeTotal must be equal 8 hr."))
	}
}

func TestManageWorkWorkingTimeMustBePresentOrFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	test_data := time.Now().Add(-12 * time.Hour)

	manageWorkTime := ManageWorkTime{
		Comment:     "ตารางงานพนักงาน",
		WorkingDate: test_data, // ผิด
		TimeTotal:   8,
		Manager:     managerChut,
		Employee:    managerChut,
		Day: Day{
			DayNumber: 26,
		},
		Month: Month{
			MonthOfYear: "มกราคม",
		},
		WorkingTime: WorkingTime{
			TimeToTime: "08.00-17.00",
		},
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(manageWorkTime)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("WorkingDate must be present or future"))
}
