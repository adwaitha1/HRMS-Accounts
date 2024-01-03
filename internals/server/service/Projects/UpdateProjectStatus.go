package Projects

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProjectStatus(c *gin.Context) {
	db := database.DB

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	var pr Projects
	if err := c.BindJSON(&pr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row, err := db.Exec(`
		UPDATE project_details
		SET proj_status=?, updated_date=NOW()
		WHERE proj_id=?
	`, pr.ProjStatus, pr.ProjID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project status" + err.Error()})
		log.Println("Error while updating project status", err)
		return
	}
	fmt.Println(row.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Project status updated"})

}
