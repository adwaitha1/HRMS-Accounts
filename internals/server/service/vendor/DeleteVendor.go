package vendor

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteVendor(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		log.Println("No DB connection!")
		return
	}
	//defer db.Close()

	Vendor_Id := c.Param("vendor_id")
	if Vendor_Id == " " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vendor ID is required"})
		return
	}

	// Check if the employee with the given ID exists
	if !vendorExists(Vendor_Id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vendor not found"})
		log.Println("Error: No such Vendor found!")
		return
	}

	rows, err := db.Exec(`DELETE FROM vendor_details WHERE vendor_id=?`, Vendor_Id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vendor"})
		log.Println("Error while deleting vendor details", err)
		return
	}
	fmt.Println(rows.RowsAffected())
	c.JSON(http.StatusOK, gin.H{"message": "Vendor deleted successfully"})
}

func vendorExists(Vendor_Id string) bool {
	db := database.DB
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM vendor_details WHERE vendor_id = ?)", Vendor_Id).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking vendor existence:", err)
		return false
	}
	return exists
}
