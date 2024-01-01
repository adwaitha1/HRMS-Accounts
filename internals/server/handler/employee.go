package handler

import (
	"HRMS/internals/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getEmployeeHandler(ctx *gin.Context) {
	log := logger.GetLogger(ctx)
	log.Info("Got a request to get Employee List")
	resp, err := h.actions.GetEmployeesDetails(ctx)
	if err != nil {
		statusCode, errorMessage := customError.HandleError(err)
		ctx.JSON(statusCode, gin.H{"error": errorMessage})
	}
	ctx.JSON(http.StatusOK, resp)
}
func (h *Handler) addNewEmployeeHandler(ctx *gin.Context) {
	log := logger.GetLogger(ctx)
	log.Info("Got a request to Create Employee")
	h.actions.AddNewEmployee(ctx)
	return
}
