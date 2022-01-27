package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team06/backend/entity"
)

func CreateMangeWorkTime(c *gin.Context) {
	var manageworktimes entity.ManageWorkTime
	var employees entity.Employee
	var days entity.Day
	var weeklies entity.Weekly
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

	if tx := entity.DB().Where("id = ?", manageworktimes.WeeklyID).First(&weeklies); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "weekly not found"})
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

	mw := entity.ManageWorkTime{
		NameSchedule: manageworktimes.NameSchedule,
		WorkingDate:  manageworktimes.WorkingDate,
		TimeTotal:    manageworktimes.TimeTotal,
		Employee:     employees,
		Day:          days,
		Weekly:       weeklies,
		Month:        months,
		WorkingTime:  worktimes,
	}

	if err := entity.DB().Create(&mw).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mw})
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
func GetAllManageWorkTime(c *gin.Context) {
	var managework []entity.ManageWorkTime
	if err := entity.DB().
		Preload("Employee.UserDetail").
		Preload("Day").
		Preload("Month").
		Preload("Weekly").
		Preload("WorkingTime").
		Raw("SELECT * FROM manage_work_times").Find(&managework).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": managework})
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
func GetAllWeekly(c *gin.Context) {
	var weeklys []entity.Weekly
	if err := entity.DB().Raw("SELECT * FROM weeklies").Scan(&weeklys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": weeklys})
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
