package Projects

import (
	"HRMS/internals/server/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Projects struct {
	ProjID        string    `json:"project_id"`
	ProjName      string    `json:"project_name"`
	ProjDesc      string    `json:"project_desc"`
	ProjStatus    string    `json:"project_status"`
	ClientID      string    `json:"client_id"`
	ResourceCount int       `json:"resource_count"`
	Vendor_Id     string    `json:"vendor_id"`
	CreatedDate   time.Time `json:"created_date"`
}

func AddProjectDetails(c *gin.Context) {
	db := database.DB
	var proj Projects
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	if err := c.BindJSON(&proj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStatement := `INSERT INTO project_details (proj_id, proj_name, proj_description, proj_status, client_id,resources_count,vendor_id,created_date)
	 	SELECT CONCAT('PR', LPAD(COALESCE(MAX(CAST(SUBSTRING(proj_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), ?, ?, ?, ?, ?, ?, NOW()
	 	FROM project_details`
	fmt.Println("SQL Statement:", sqlStatement)

	var rows sql.Result

	rows, err := db.Exec(sqlStatement, proj.ProjName, proj.ProjDesc, proj.ProjStatus, proj.ClientID, proj.ResourceCount, proj.Vendor_Id)
	fmt.Printf("VALUES:%s,%s,%s,%s,%d,%s\n", proj.ProjName, proj.ProjDesc, proj.ProjStatus, proj.ClientID, proj.ResourceCount, proj.Vendor_Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting employee"})
		log.Println("Error inserting project details:", err)
		return
	}
	fmt.Println(rows.RowsAffected())
	var lastInsertID string
	err = db.QueryRow("SELECT CONCAT('PR', LPAD(COALESCE(MAX(CAST(SUBSTRING(proj_id, 3) AS UNSIGNED)), 0), 3, '0')) FROM project_details").Scan(&lastInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving last insert ID", "details": err.Error()})
		log.Println("Error retrieving last insert ID:", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("New Project added successfully\nProject ID: %v", lastInsertID),
	})
	fmt.Println("Project details inserted")

}
