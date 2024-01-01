package service

import (
	"HRMS/internals/pkg/logger"
	m "HRMS/internals/server/Models"
	"HRMS/internals/server/customError"

	"github.com/gin-gonic/gin"
)

// GetEmployeesDetails fetches list of employees
func (a *Actions) GetEmployeesDetails(c *gin.Context) ([]m.Employees, error) {
	log := logger.GetLogger(c)
	log.Info("GetEmployeesDetails: Processing Get Employee Details")
	emp, err := a.db.GetEmployeesDetails(c)
	if err != nil {
		log.Errorln("GetEmployeesDetails: Error while fetching employee details Error:", err.Error())
		return nil, &customError.DatabaseError{Reason: "GetEmployeesDetails: Error while fetching employee details Error:" + err.Error()}

	}
	return emp, err
}

// AddNewEmployee will create employee information
func (a *Actions) AddNewEmployee(c *gin.Context) (string, error) {
	log := logger.GetLogger(c)
	log.Info("AddNewEmployee: Processing New Employee information")
	var emp m.Employees
	if err := c.BindJSON(&emp); err != nil {
		log.Errorln("AddNewEmployee: Error while binding request information Error:", err.Error())
		return "", &customError.ServiceError{Reason: "AddNewEmployee: Error while binding request information Error:" + err.Error()}
	}
	lastInsertId, err := a.db.AddNewEmployee(c, emp)
	if err != nil {
		log.Errorln("AddNewEmployee: Error while creating employee information Error:", err.Error())
		return "", &customError.DatabaseError{Reason: "AddNewEmployee: Error while creating employee information Error:" + err.Error()}
	}
	return lastInsertId, nil
}

// UpdateEmployeeDetails updates leave_count and working_hours in the employees_details table
func (a *Actions) UpdateEmployeeLeave(c *gin.Context) error {
	log := logger.GetLogger(c)
	var emp m.EmployeeDetails
	if err := c.BindJSON(&emp); err != nil {
		log.Errorln("UpdateEmployeeLeave: Error while binding request information Error:", err.Error())
		return &customError.ServiceError{Reason: "UpdateEmployeeLeave: Error while binding request information Error:" + err.Error()}
	}
	if emp.LeaveCount != 0 {
		emp.WorkingHours = float64(TotalWorkingDays-int(emp.LeaveCount)) * 8.0

	} else {
		emp.WorkingHours = float64(TotalWorkingDays) * 8.0
	}
	err := a.db.UpdateEmployeeLeave(c, emp)
	if err != nil {
		log.Errorln("UpdateEmployeeLeave: Error while updating employee information, Error:", err.Error())
		return &customError.DatabaseError{Reason: "UpdateEmployeeLeave: Error while updating employee information, Error:" + err.Error()}
	}
	return nil
}

// func DeleteEmployee(c *gin.Context) {
// 	db := database.DB
// 	if db == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
// 		log.Println("No DB connection!")
// 		return
// 	}
// 	//defer db.Close()

// 	empID := c.Param("emp_id")
// 	if empID == " " {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee ID is required"})
// 		return
// 	}

// 	// Check if the employee with the given ID exists
// 	if !employeeExists(empID) {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
// 		log.Println("Error: No such Employee found!")
// 		return
// 	}

// 	rows, err := db.Exec(`DELETE FROM employees_details WHERE emp_id=?`, empID)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
// 		log.Println("Error while deleting employee details", err)
// 		return
// 	}
// 	fmt.Println(rows.RowsAffected())
// 	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
// }

// func employeeExists(empID string) bool {
// 	db := database.DB
// 	var exists bool
// 	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM employees_details WHERE emp_id = ?)", empID).Scan(&exists)
// 	if err != nil {
// 		fmt.Println("Error checking employee existence:", err)
// 		return false
// 	}
// 	return exists
// }
