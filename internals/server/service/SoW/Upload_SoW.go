package SoW

import (
	"HRMS/internals/server/database"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type SowDetails struct {
	SowID                 string    `db:"sow_id"`
	SowFilePath           string    `db:"sow_filepath"`
	ProjID                string    `db:"proj_id"`
	IsInbound             bool      `db:"isinbound"`
	StartDate             time.Time `db:"start_date"`
	EndDate               time.Time `db:"end_date"`
	ClientID              string    `db:"client_id"`
	AgreementDate         time.Time `db:"agreement_date"`
	AgreementExecDate     time.Time `db:"agreement_exec_date"`
	SowTermsAgreed        bool      `db:"sow_terms_agreed"`
	Amount                float64   `db:"amount"`
	BillingFrequency      string    `db:"billing_frequency"`
	SupplierSignaturePath string    `db:"supplier_signature_path"`
	CreatedDate           time.Time `db:"created_date"`
	UpdatedDate           time.Time `db:"updated_date"`
	VendorID              string    `db:"vendor_id"`
	Products              string    `db:"products"`
}

func FileUpload(c *gin.Context) (string, error) {
	file, err := c.FormFile("File")
	if err != nil {
		fmt.Println("Error retrieving the file:", err)
		return "", err
	}

	filePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		fmt.Println("Error saving the file:", err)
		return "", err
	}

	return filePath, nil
}

func InsertSowDetails(c *gin.Context) {
	db := database.DB
	var Sow SowDetails
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}

	body, err := c.GetRawData()
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	fmt.Println(string(body))
	if err := c.ShouldBind(&Sow); err != nil {
		fmt.Println("Error while Binding request data", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the VALUES clause directly for generating sow_id
	rows, err := db.Exec(`INSERT INTO sow_details
		(sow_id, sow_filepath, proj_id, isinbound, start_date, end_date, client_id, agreement_date, 
			agreement_exec_date, sow_terms_agreed, amount, billing_frequency, supplier_signature_path, created_date, vendor_id)
		SELECT CONCAT('SOW', LPAD(COALESCE(MAX(CAST(SUBSTRING(sow_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), 
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), ? FROM sow_details`, Sow.SowFilePath, Sow.ProjID, Sow.IsInbound, Sow.StartDate, Sow.EndDate, Sow.ClientID, Sow.AgreementDate, Sow.AgreementExecDate, Sow.SowTermsAgreed, Sow.Amount, Sow.BillingFrequency, Sow.SupplierSignaturePath, Sow.VendorID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database"})
		log.Println("Error while Uploading SoW:", err)
		return
	}
	fmt.Println(rows.RowsAffected())

	c.JSON(http.StatusOK, gin.H{"message": "Data inserted successfully"})
}
