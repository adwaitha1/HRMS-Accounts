package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 'Route' defines the structure of a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

// 'NewRouter' creates a new Gin router with the provided routes and middleware
func NewRouter(custom Routes, middleware ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	for _, m := range middleware {
		router.Use(m)
	}

	// Add routes
	for _, route := range custom {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		default:
			panic("Unsupported HTTP method")
		}
	}

	return router
}
