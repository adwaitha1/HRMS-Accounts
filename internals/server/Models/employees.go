package models

import "time"

type Employees struct {
	EmpID          string  `json:"employee_id"`
	EmployeeName   string  `json:"employee_name"`
	EmployeeRole   string  `json:"employee_role"`
	EmployeeSkills string  `json:"employee_skills"`
	Experience     float32 `json:"experience"`
}
type EmployeeDetails struct {
	EmpID        string    `json:"emp_id"`
	LeaveCount   float64   `json:"leave_count"`
	Months       string    `json:"months"`
	WorkingHours float64   `json:"working_hours"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}
