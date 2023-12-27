package Client

import (
	"HRMS/internals/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClientDetails(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	query := "SELECT client_id, client_company_name, client_name, client_address1, client_address2,business_entity_type, country, currency, state, state_code, gst_rates FROM client_details WHERE is_active=1"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var client []Client
	for rows.Next() {
		var cli Client
		err := rows.Scan(
			&cli.ClientID,
			&cli.ClientCompany,
			&cli.ClientName,
			&cli.ClientAddr1,
			&cli.ClientAddr2,
			&cli.BusinessEntityType,
			&cli.Country,
			&cli.Currency,
			&cli.State,
			&cli.State_code,
			&cli.Gst_rate,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		client = append(client, cli)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}
