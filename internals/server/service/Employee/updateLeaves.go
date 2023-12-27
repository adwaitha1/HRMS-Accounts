package Employee

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EmployeeDetails struct {
	EmpID        string    `json:"emp_id"`
	LeaveCount   float64   `json:"leave_count"`
	Months       string    `json:"months"`
	WorkingHours float64   `json:"working_hours"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

// UpdateEmployeeDetails updates leave_count and working_hours in the employees_details table
func UpdateEmployeeLeave(c *gin.Context) {
	db := database.DB

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	var emp EmployeeDetails
	if err := c.BindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	totalWorkingDays := 22
	if emp.LeaveCount != 0 {
		emp.WorkingHours = float64(totalWorkingDays-int(emp.LeaveCount)) * 8.0

	} else {
		emp.WorkingHours = float64(totalWorkingDays) * 8.0
	}

	row, err := db.Exec(`
		UPDATE employees_details
		SET leave_count=?,months=?, working_hours=?, updated_date=NOW()
		WHERE emp_id=?
	`, emp.LeaveCount, emp.Months, emp.WorkingHours, emp.EmpID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee details" + err.Error()})
		log.Println("Error while updating employee details", err)
		return
	}
	fmt.Println(row.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Employee details updated successfully"})
}
