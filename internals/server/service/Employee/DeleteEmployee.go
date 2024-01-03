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
	// if !employeeExists(empID) {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	// 	log.Println("Error: No such Employee found!")
	// 	return
	// }
	if !employeeExists(empID) {
		log.Println("ERROR: No such Employee found!")
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Employee not found", "success": false})
		return
	}

	rows, err := db.Exec(`DELETE FROM employees_details WHERE emp_id=?`, empID)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
	// 	log.Println("Error while deleting employee details", err)
	// 	return
	// }
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Failed to delete employee: " + err.Error(), "success": false})
		return
	}

	fmt.Println(rows.RowsAffected())
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Employee deleted successfully",
		"success": true,
	})
}

func employeeExists(empID string) bool {
	db := database.DB
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM employees_details WHERE emp_id = ?)", empID).Scan(&exists)
	if err != nil {
		log.Println("ERROR:", err)
		return false
	}

	return exists
}
