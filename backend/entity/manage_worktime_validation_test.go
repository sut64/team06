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

func TestManageWorkTimeTimeToTalMustEqualEight(t *testing.T) {
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
