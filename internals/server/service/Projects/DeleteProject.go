package Projects

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteProject(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		log.Println("No DB connection!")
		return
	}
	//defer db.Close()

	ProjectID := c.Param("proj_id")
	if ProjectID == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project ID is required"})
		return
	}

	// Check if the project with the given ID exists
	if !projectExists(ProjectID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		log.Println("Error: project does not exist!")
		return
	}

	rows, err := db.Exec(`DELETE FROM project_details WHERE proj_id=?`, ProjectID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project details"})
		log.Println("Error while deleting a project details", err)
		return
	}
	fmt.Println(rows.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Project removed successfully"})
}

func projectExists(ProjectID string) bool {
	db := database.DB
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM project_details WHERE proj_id = ?)", ProjectID).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking the project existence:", err)
		return false
	}
	return exists
}
