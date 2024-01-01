package service

import (
	m "HRMS/internals/server/Models"

	"github.com/gin-gonic/gin"
)

type DB interface {
	GetEmployeesDetails(c *gin.Context) ([]m.Employees, error)
	AddNewEmployee(c *gin.Context, emp m.Employees) (string, error)
	UpdateEmployeeLeave(c *gin.Context, emp m.EmployeeDetails) error
}

type Actions struct {
	DB DB
}

func NewActions(db DB) *Actions {
	return &Actions{
		DB: db,
	}
}

const (
	TotalWorkingDays int = 22
)
