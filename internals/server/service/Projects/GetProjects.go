package Projects

import (
	"HRMS/internals/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProjectDetails(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	query := "SELECT proj_id, proj_name, proj_description, proj_status, client_id,resources_count,vendor_id FROM project_details"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var project []Projects
	for rows.Next() {
		var proj Projects
		err := rows.Scan(
			&proj.ProjID,
			&proj.ProjName,
			&proj.ProjDesc,
			&proj.ProjStatus,
			&proj.ClientID,
			&proj.ResourceCount,
			&proj.Vendor_Id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		project = append(project, proj)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, project)
}
