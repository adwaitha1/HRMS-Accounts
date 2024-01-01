package handler

import (
	m "HRMS/internals/server/Models"
	"HRMS/internals/server/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	actions actions
}

type actions interface {
	GetEmployeesDetails(c *gin.Context) ([]m.Employees, error)
	AddNewEmployee(c *gin.Context) (string, error)
	UpdateEmployeeLeave(c *gin.Context) error
}

func NewHandler(actions actions) *Handler {
	return &Handler{
		actions: actions,
	}
}

const (
	employeePath = "/employees"
	vendorPath   = "/vendors"
	clientPath   = "/clients"
	projectsPath = "/projects"
	sowPath      = "/sow"
	idPath       = "/:id"
)

func (h *Handler) Routes() routers.Routes {
	return []routers.Route{
		{
			Name:        "Get Employee List",
			Method:      http.MethodGet,
			Pattern:     employeePath,
			HandlerFunc: h.getEmployeeHandler,
		},
	}
}

// router.GET("/Employee_List", emp.GetEmployeesDetails)
// router.POST("/addNewEmployee", emp.AddNewEmployee)
// router.PATCH("/update_Employee_Leave", emp.UpdateEmployeeLeave)
// router.DELETE("/delete-employee/:emp_id", emp.DeleteEmployee)
// router.POST("/addNewVendor", vendor.AddNewVendor)
// router.GET("/Vendors_List", vendor.GetVendorDetails)
// router.DELETE("/delete_vendor/:vendor_id", vendor.DeleteVendor)
// router.POST("/addNewClient", Client.AddNewClient)
// router.GET("/viewClients", Client.GetClientDetails)
// router.DELETE("/delete-client/:client_id", Client.DeleteClient)
// router.POST("/addNewSoW", SoW.InsertSowDetails)
// router.GET("/GetProjects", pr.GetProjectDetails)
// router.POST("/addNewProject", pr.AddProjectDetails)
// router.DELETE("/delete-project/:proj_id", pr.DeleteProject)
