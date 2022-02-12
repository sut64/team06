package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team06/backend/entity"
)

func CreateMangeWorkTime(c *gin.Context) {
	var manageworktimes entity.ManageWorkTime
	var employees entity.Employee
	var days entity.Day
	var months entity.Month
	var worktimes entity.WorkingTime

	if err := c.ShouldBindJSON(&manageworktimes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", manageworktimes.EmployeeID).First(&employees); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", manageworktimes.DayID).First(&days); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "day not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", manageworktimes.MonthID).First(&months); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "month not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", manageworktimes.WorkingTimeID).First(&worktimes); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "working time not found"})
		return
	}

	mwt := entity.ManageWorkTime{
		Comment:     manageworktimes.Comment,
		WorkingDate: manageworktimes.WorkingDate,
		TimeTotal:   manageworktimes.TimeTotal,
		Employee:    employees,
		Manager:     employees,
		Day:         days,
		Month:       months,
		WorkingTime: worktimes,
	}

	if _, err := govalidator.ValidateStruct(mwt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&mwt).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mwt})
}

func GetAllManageWorkTime(c *gin.Context) {
	var managework []entity.ManageWorkTime
	if err := entity.DB().
		Preload("Employee.UserDetail").
		Preload("Employee.Position").
		Preload("Day").
		Preload("Month").
		Preload("WorkingTime").
		Raw("SELECT * FROM manage_work_times").Find(&managework).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": managework})
}

func DeleteManageWorkTime(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM manage_work_times WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MangeWorkTime not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

func GetAllDetail(c *gin.Context) {
	var details []entity.UserDetail
	if err := entity.DB().Raw("SELECT * FROM user_details").Scan(&details).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": details})
}

func GetAllPosition(c *gin.Context) {
	var positions []entity.EmployeePosition
	if err := entity.DB().Raw("SELECT * FROM employee_positions").Scan(&positions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": positions})
}

func GetAllEmployee(c *gin.Context) {
	var employees []entity.Employee
	if err := entity.DB().
		Preload("UserDetail").
		Preload("Position").
		Preload("UserLogin").
		Raw("SELECT * FROM employees").Find(&employees).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func GetAllDay(c *gin.Context) {
	var days []entity.Day
	if err := entity.DB().Raw("SELECT * FROM days").Scan(&days).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": days})
}

func GetAllMonth(c *gin.Context) {
	var month []entity.Month
	if err := entity.DB().Raw("SELECT * FROM months").Scan(&month).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": month})
}
func GetAllWorkTime(c *gin.Context) {
	var worktime []entity.WorkingTime
	if err := entity.DB().Raw("SELECT * FROM working_times").Scan(&worktime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": worktime})
}
