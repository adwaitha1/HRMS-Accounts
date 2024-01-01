package Employee

import (
	"HRMS/internals/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employees struct {
	EmpID          string  `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	EmployeeRole   string  `json:"employee_role"`
	EmployeeSkills string  `json:"employee_skills"`
	Experience     float32 `json:"experience"`
}

func GetEmployeesDetails(c *gin.Context) {

	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	query := "SELECT emp_id, emp_name, emp_role, emp_skills, experience FROM employees_details"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var employees []Employees
	for rows.Next() {
		var employee Employees
		err := rows.Scan(
			&employee.EmpID,
			&employee.EmployeeName,
			&employee.EmployeeRole,
			&employee.EmployeeSkills,
			&employee.Experience,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}
