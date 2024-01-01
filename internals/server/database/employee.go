package database

import (
	"HRMS/internals/pkg/logger"
	m "HRMS/internals/server/Models"

	"github.com/gin-gonic/gin"
)

func (db *Database) GetEmployeesDetails(c *gin.Context) ([]m.Employees, error) {
	log := logger.GetLogger(c)
	query := "SELECT emp_id, emp_name, emp_role, emp_skills, experience FROM employees_details"
	rows, err := db.dbClient.Query(query)
	if err != nil {
		log.Errorln("GetEmployeesDetails: Failed to query employee details Error:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var employees []m.Employees
	for rows.Next() {
		var employee m.Employees
		err := rows.Scan(
			&employee.EmpID,
			&employee.EmployeeName,
			&employee.EmployeeRole,
			&employee.EmployeeSkills,
			&employee.Experience,
		)
		if err != nil {
			log.Errorln("GetEmployeesDetails: Failed to scan all employee details Error:", err.Error())
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Errorln("GetEmployeesDetails: Failed to scan all employee details Error:", err.Error())
		return nil, err
	}
	log.Infoln("GetEmployeesDetails: Fetched all employee details, Count:", len(employees))
	return employees, nil
}
func (db *Database) AddNewEmployee(c *gin.Context, emp m.Employees) (string, error) {
	log := logger.GetLogger(c)

	sqlStatement := `INSERT INTO employees_details (emp_id, emp_name, emp_role, emp_skills, experience, created_date)
	 	SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(emp_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), ?, ?, ?, ?, NOW()
	 	FROM employees_details`
	//fmt.Println("SQL Statement:", sqlStatement)

	rows, err := db.dbClient.Exec(sqlStatement, emp.EmployeeName, emp.EmployeeRole, emp.EmployeeSkills, emp.Experience)
	//fmt.Printf("VALUES:%s,%s,%s,%f\n", emp.EmployeeName, emp.EmployeeRole, emp.EmployeeSkills, emp.Experience)
	if err != nil {
		log.Errorln("AddNewEmployee: Failed to execute insert employee, Error:", err.Error())
		return "", err
	}
	rowsAffected, _ := rows.RowsAffected()
	log.Infoln("Inserted employee information, Rows Affected :", rowsAffected)
	var lastInsertID string
	err = db.dbClient.QueryRow("SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(emp_id, 3) AS UNSIGNED)), 0), 3, '0')) FROM employees_details").Scan(&lastInsertID)
	if err != nil {
		log.Errorln("AddNewEmployee: Error retrieving last insert ID:", err)
		return "", err
	}
	return lastInsertID, err
}

// UpdateEmployeeDetails updates leave count and working hours in the employees_details table
func (db *Database) UpdateEmployeeLeave(c *gin.Context, emp m.EmployeeDetails) error {
	log := logger.GetLogger(c)
	row, err := db.dbClient.Exec(`
		UPDATE employees_details
		SET leave_count=?,months=?, working_hours=?, updated_date=NOW()
		WHERE emp_id=?
	`, emp.LeaveCount, emp.Months, emp.WorkingHours, emp.EmpID)

	if err != nil {
		log.Errorln("Error while updating employee details", err)
		return err
	}
	rowsAffected, _ := row.RowsAffected()
	log.Infoln("Employee details updated successfully, Rows Affected :", rowsAffected)
	return nil
}
