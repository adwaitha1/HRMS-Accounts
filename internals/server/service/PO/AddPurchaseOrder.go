package po

import (
	"HRMS/internals/server/database"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PurchaseOrder struct {
	PO_Id         string    `json:"po_Id"`
	PO_No         string    `json:"po_No"`
	Rev_No        string    `json:"rev_No"`
	PO_Date       time.Time `json:"po_Date"`
	Delivery_Date time.Time `json:"delivery_Date"`
}

func AddNewProject(c *gin.Context) {
	db := database.DB
	var po PurchaseOrder
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	if err := c.BindJSON(&po); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStatement := `INSERT INTO employees_details (emp_id, emp_name, emp_role, emp_skills, experience, created_date)
	SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(emp_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), ?, ?, ?, ?, NOW()
	FROM employees_details`
	fmt.Println("SQL Statement:", sqlStatement)

}
