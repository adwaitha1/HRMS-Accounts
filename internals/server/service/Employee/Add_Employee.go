package Employee

import (
	"HRMS/internals/server/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddNewEmployee(c *gin.Context) {
	db := database.DB
	var emp Employees
	// if db == nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
	// 	return
	// }
	if db == nil {
		log.Println("ERROR: Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Database connection is nil", "success": false})
		return
	}
	// if err := c.BindJSON(&emp); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := c.BindJSON(&emp); err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Bad Request", "showMessage": "Developer", "details": err.Error(), "success": false})
		return
	}
	sqlStatement := `INSERT INTO employees_details (emp_id, emp_name, emp_role, emp_skills, experience, created_date)
	 	SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(emp_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), ?, ?, ?, ?, NOW()
	 	FROM employees_details`
	fmt.Println("SQL Statement:", sqlStatement)

	var rows sql.Result

	rows, err := db.Exec(sqlStatement, emp.EmployeeName, emp.EmployeeRole, emp.EmployeeSkills, emp.Experience)
	//fmt.Printf("VALUES:%s,%s,%s,%f\n", emp.EmployeeName, emp.EmployeeRole, emp.EmployeeSkills, emp.Experience)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting employee"})
	// 	log.Println("Error inserting employee:", err)
	// 	return
	// }
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Error inserting employee: " + err.Error(), "success": false})
		return
	}
	fmt.Println(rows.RowsAffected())
	var lastInsertID string
	err = db.QueryRow("SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(emp_id, 3) AS UNSIGNED)), 0), 3, '0')) FROM employees_details").Scan(&lastInsertID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving last insert ID", "details": err.Error()})
	// 	log.Println("Error retrieving last insert ID:", err)
	// 	return
	// }
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Error retrieving last insert ID: " + err.Error(), "success": false})
		return
	}

	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": fmt.Sprintf("Employee added successfully\nEmployee ID: %v", lastInsertID),
	// })
	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Employee created successfully",
		"data":    map[string]interface{}{"employeeID": lastInsertID},
		"success": true,
	})

	fmt.Println("Employee data added successfully.")
}
