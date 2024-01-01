package Projects

import (
	"HRMS/internals/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProjectStatus(c *gin.Context) {
	db := database.DB

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	var status Projects
	if err := c.BindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if status.ProjStatus!="completed"
}
