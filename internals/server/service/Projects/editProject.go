package Projects

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// adding the line
// to git check
func UpdateProjectCount(c *gin.Context) {
	db := database.DB

	if db == nil {
		log.Println("ERROR: Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Database connection is nil", "success": false})
		return
	}
	var pr Projects
	// if err := c.BindJSON(&pr); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err := c.BindJSON(&pr); err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Bad Request", "showMessage": "Developer", "details": err.Error(), "success": false})
		return
	}

	row, err := db.Exec(`
		UPDATE project_details
		SET resources_count=?, updated_date=NOW()
		WHERE proj_id=?
	`, pr.ProjStatus, pr.ProjID)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update resource count" + err.Error()})
	// 	log.Println("Error while updating resources count", err)
	// 	return
	// }
	if err != nil {
		log.Println("ERROR:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal Server Error", "showMessage": "Developer", "details": "Failed to update resource count: " + err.Error(), "success": false})
		return
	}

	fmt.Println(row.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Resource count changed", "success": true})
}
