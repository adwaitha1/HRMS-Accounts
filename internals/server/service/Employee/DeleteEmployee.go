package Employee

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEmployee(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		log.Println("No DB connection!")
		return
	}
	//defer db.Close()

	empID := c.Param("emp_id")
	if empID == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee ID is required"})
		return
	}

	// Check if the employee with the given ID exists
	if !employeeExists(empID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		log.Println("Error: No such Employee found!")
		return
	}

	rows, err := db.Exec(`DELETE FROM employees_details WHERE emp_id=?`, empID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		log.Println("Error while deleting employee details", err)
		return
	}
	fmt.Println(rows.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

func employeeExists(empID string) bool {
	db := database.DB
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM employees_details WHERE emp_id = ?)", empID).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking employee existence:", err)
		return false
	}
	return exists
}
