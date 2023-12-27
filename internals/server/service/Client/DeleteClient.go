package Client

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteClient(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		log.Println("No DB connection!")
		return
	}
	//defer db.Close()

	ClientID := c.Param("client_id")
	if ClientID == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Client ID is required"})
		return
	}

	// Check if the employee with the given ID exists
	if !employeeExists(ClientID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		log.Println("Error: No such client found!")
		return
	}

	rows, err := db.Exec(`DELETE FROM Client_details WHERE client_id=?`, ClientID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		log.Println("Error while deleting client details", err)
		return
	}
	fmt.Println(rows.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Client deleted successfully"})
}

func employeeExists(ClientID string) bool {
	db := database.DB
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM client_details WHERE client_id = ?)", ClientID).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking client existence:", err)
		return false
	}
	return exists
}
