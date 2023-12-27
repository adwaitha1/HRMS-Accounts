package vendor

import (
	"HRMS/internals/server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVendorDetails(c *gin.Context) {
	db := database.DB
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
		return
	}
	query := "SELECT vendor_id,vendor_no,vendor_company, vendor_name, vendor_ph_no, vendor_email, vendor_addr1,vendor_addr2,vendor_city,vendor_state,vendor_state_code,country,pincode,vendor_pan_no,vendor_gst_no,vendor_tan_no,vendor_bank_name,acc_holder_name,bank_acc_no,vendor_ifsc_code,onboarding_date,offboarding_date,status_type FROM vendor_details"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var vendors []Vendors
	for rows.Next() {
		var vendor Vendors
		err := rows.Scan(
			&vendor.Vendor_Id,
			&vendor.Vendor_No,
			&vendor.Vendor_company,
			&vendor.Vendor_name,
			&vendor.Vendor_Ph_No,
			&vendor.Vendor_Email,
			&vendor.Vendor_Addr1,
			&vendor.Vendor_Addr2,
			&vendor.Vendor_City,
			&vendor.State,
			&vendor.StateCode,
			&vendor.Country,
			&vendor.Pincode,
			&vendor.Vendor_PAN_No,
			&vendor.Vendor_GST_No,
			&vendor.Vendor_TAN_No,
			&vendor.Vendor_BankName,
			&vendor.Vendor_AC_holder_name,
			&vendor.Vendor_Bank_AC_No,
			&vendor.Vendor_Bank_IFSC_Code,
			&vendor.OnBoarding_Date,
			&vendor.OffBoarding_Date,
			&vendor.Status_Type,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		vendors = append(vendors, vendor)
	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vendors)
}
