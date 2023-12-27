// vendor.go
package vendor

import (
	"HRMS/internals/server/database"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Vendors struct {
	Vendor_Id             string `json:"vendor_id"`
	Vendor_No             string `json:"vendor_no"`
	Vendor_company        string `json:"vendor_company"`
	Vendor_name           string `json:"vendor_name"`
	Vendor_Ph_No          string `json:"vendor_ph_no"`
	Vendor_Email          string `json:"vendor_email"`
	Vendor_Addr1          string `json:"vendor_addr1"`
	Vendor_Addr2          string `json:"vendor_addr2"`
	Vendor_City           string `json:"vendor_city"`
	State                 string `json:"state"`
	StateCode             int    `json:"state_code"`
	Country               string `json:"country"`
	Pincode               string `json:"pincode"`
	Vendor_PAN_No         string `json:"vendor_pan_no"`
	Vendor_GST_No         string `json:"vendor_gst_no"`
	Vendor_TAN_No         string `json:"vendor_tan_no"`
	Vendor_BankName       string `json:"vendor_bank_name"`
	Vendor_AC_holder_name string `json:"vendor_ac_holder_name"`
	Vendor_Bank_AC_No     string `json:"vendor_bank_ac_no"`
	Vendor_Bank_IFSC_Code string `json:"vendor_bank_ifsc_code"`
	OnBoarding_Date       string `json:"onboarding_date"`
	OffBoarding_Date      string `json:"offboarding_date"`
	Status_Type           string `json:"status_type"`
	IsActive              int    `json:"isactive"`
}

func handleError(c *gin.Context, statusCode int, errorMsg string, err error) {
	c.JSON(statusCode, gin.H{"error": errorMsg, "details": err.Error()})
	c.Error(err).SetType(gin.ErrorTypePrivate)
}

func parseDate(dateString string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, dateString)
}

func executeQuery(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

func getLastInsertID(db *sql.DB) (string, error) {
	var lastInsertID string
	err := db.QueryRow("SELECT CONCAT('IB', LPAD(COALESCE(MAX(CAST(SUBSTRING(vendor_id, 3) AS UNSIGNED)), 0), 3, '0')) FROM vendor_details").Scan(&lastInsertID)
	return lastInsertID, err
}

func AddNewVendor(c *gin.Context) {
	db := database.DB
	if db == nil {
		handleError(c, http.StatusInternalServerError, "Database connection is nil", nil)
		return
	}

	var vendor Vendors
	if err := c.ShouldBindJSON(&vendor); err != nil {
		handleError(c, http.StatusBadRequest, "Error binding JSON", err)
		return
	}

	if vendor.OnBoarding_Date == "" || vendor.OffBoarding_Date == "" {
		handleError(c, http.StatusBadRequest, "Date field is empty", nil)
		return
	}

	onBoardingDate, err := parseDate(vendor.OnBoarding_Date)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Error parsing onboarding date", err)
		return
	}

	offBoardingDate, err := parseDate(vendor.OffBoarding_Date)
	if err != nil {
		handleError(c, http.StatusBadRequest, "Error parsing offboarding date", err)
		return
	}

	query := `
		INSERT INTO vendor_details (vendor_id, vendor_no, vendor_company, vendor_name, vendor_ph_no, vendor_email,
			vendor_addr1, vendor_addr2, vendor_city, vendor_state, vendor_state_code, country, pincode, 
			vendor_pan_no, vendor_gst_no, vendor_tan_no, vendor_bank_name, acc_holder_name, bank_acc_no,
			 vendor_ifsc_code, onboarding_date, offboarding_date, status_type, is_active, created_date)
		SELECT CONCAT('VN', LPAD(COALESCE(MAX(CAST(SUBSTRING(vendor_id, 3) AS UNSIGNED)), 0) + 1, 3, '0')), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW()
		FROM vendor_details
	`

	result, err := executeQuery(db, query,
		vendor.Vendor_No, vendor.Vendor_company, vendor.Vendor_name, vendor.Vendor_Ph_No, vendor.Vendor_Email,
		vendor.Vendor_Addr1, vendor.Vendor_Addr2, vendor.Vendor_City, vendor.State, vendor.StateCode,
		vendor.Country, vendor.Pincode, vendor.Vendor_PAN_No, vendor.Vendor_GST_No,
		vendor.Vendor_TAN_No, vendor.Vendor_BankName, vendor.Vendor_AC_holder_name, vendor.Vendor_Bank_AC_No,
		vendor.Vendor_Bank_IFSC_Code, onBoardingDate, offBoardingDate, vendor.Status_Type, vendor.IsActive,
	)

	if err != nil {
		handleError(c, http.StatusInternalServerError, "Error inserting Vendor", err)
		return
	}

	fmt.Println(result.RowsAffected())

	lastInsertID, err := getLastInsertID(db)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Error retrieving last insert ID", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Vendor added successfully\nVendor ID: %v", lastInsertID),
	})

	fmt.Println("Vendor details added successfully.")
}
