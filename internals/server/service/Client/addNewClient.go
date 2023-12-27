package Client

import (
	"HRMS/internals/server/database"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	ClientID           string    `json:"clientId"`
	ClientCompany      string    `json:"client_company"`
	ClientName         string    `json:"client_name"`
	ClientAddr1        string    `json:"client_addr1"`
	ClientAddr2        string    `json:"client_addr2"`
	BusinessEntityType string    `json:"business_entity_type"`
	Country            string    `json:"country"`
	Currency           string    `json:"currency"`
	State              string    `json:"state"`
	State_code         int       `json:"state_code"`
	Gst_rate           string    `json:"gst_rate"`
	Created_date       time.Time `json:"created_date"`
	isActive           int       `json:"isActive"`
}

func AddNewClient(c *gin.Context) {
	db := database.DB
	var cli Client
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	if err := c.BindJSON(&cli); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStatement := `INSERT INTO client_details (
		client_id, client_company_name, client_name, client_address1, client_address2,
		business_entity_type, country, currency, state, state_code, gst_rates, created_date,is_active) 
	SELECT CONCAT('CL', LPAD(COALESCE(MAX(CAST(SUBSTRING(client_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')),
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(),1 FROM client_details`
	fmt.Println("SQL Statement:", sqlStatement)
	var rows sql.Result
	rows, err := db.Exec(sqlStatement, cli.ClientCompany, cli.ClientName, cli.ClientAddr1, cli.ClientAddr2, cli.BusinessEntityType, cli.Country, cli.Currency, cli.State, cli.State_code, cli.Gst_rate)
	fmt.Printf("VALUES:%s,%s,%s,%s,%s,%s,%s,%s,%d,%s,%d\n", cli.ClientCompany, cli.ClientName, cli.ClientAddr1, cli.ClientAddr2, cli.BusinessEntityType, cli.Country, cli.Currency, cli.State, cli.State_code, cli.Gst_rate, cli.isActive)

	if err != nil {
		fmt.Println("Error inserting Client details:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting Client details"})
		return
	}

	fmt.Println(rows.RowsAffected())

	var lastInsertID string
	err = db.QueryRow("SELECT CONCAT('CL', LPAD(COALESCE(MAX(CAST(SUBSTRING(client_id, 3) AS UNSIGNED)), 0), 3, '0')) FROM client_details").Scan(&lastInsertID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving last insert ID", "details": err.Error()})
		log.Println("Error retrieving last insert ID:", err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Client added successfully\nClient ID: %v", lastInsertID),
	})
	fmt.Println("New Client added successfully.")
}
