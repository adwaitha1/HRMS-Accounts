package main

import (
	"HRMS/internals/server/config"
	db "HRMS/internals/server/database"
	"HRMS/internals/server/service"
	"HRMS/internals/server/service/Client"
	emp "HRMS/internals/server/service/Employee"
	pr "HRMS/internals/server/service/Projects"
	"HRMS/internals/server/service/SoW"
	"HRMS/internals/server/service/vendor"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

var ServerConfig *config.ServerConfig
var err error

func init() {
	// Perform any initialization tasks here
	fmt.Println("Initializing...")
	var configPath string
	flag.StringVar(&configPath, "configpath", "", "File path for server configuration")
	flag.Parse()
	ServerConfig, err = config.LoadServerConfig(configPath)
	if err != nil {
		log.Fatal("Error loading server config : ", err)
	}
	log.Printf("Address: %s\n", ServerConfig.Database.Host)
	log.Printf("Port: %s\n", ServerConfig.Database.Port)
	log.Printf("User: %s\n", ServerConfig.Database.User)
	log.Printf("Password: %s\n", ServerConfig.Database.Password)
	log.Printf("DB Name: %s\n", ServerConfig.Database.DBName)
	log.Printf("Server Port: %d\n", ServerConfig.Server.Port)

}

func main() {

	dbConn, err := db.CreateCon(ServerConfig.Database)
	if err != nil {
		log.Fatal("Unable to Initialise DB connection ! Error:", err.Error())
	}

	serviceActions := service.NewActions(dbConn)

	// h := handler.
	fmt.Println(serviceActions)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome!!!",
		})
	})
	router.GET("/Employee_List", emp.GetEmployeesDetails)
	router.POST("/addNewEmployee", emp.AddNewEmployee)
	router.PATCH("/update_Employee_Leave", emp.UpdateEmployeeLeave)
	router.DELETE("/delete-employee/:emp_id", emp.DeleteEmployee)
	router.POST("/addNewVendor", vendor.AddNewVendor)
	router.GET("/Vendors_List", vendor.GetVendorDetails)
	router.DELETE("/delete_vendor/:vendor_id", vendor.DeleteVendor)
	router.POST("/addNewClient", Client.AddNewClient)
	router.GET("/viewClients", Client.GetClientDetails)
	router.DELETE("/delete-client/:client_id", Client.DeleteClient)
	router.POST("/addNewSoW", SoW.InsertSowDetails)
	router.GET("/GetProjects", pr.GetProjectDetails)
	router.POST("/addNewProject", pr.AddProjectDetails)
	router.DELETE("/delete-project/:proj_id", pr.DeleteProject)

	//router.POST("AddNewClient", SoW.InsertClientDetails)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Failed to start the server:", err)
		}
	}()
	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Println("Server is shutting down...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server shutdown error:", err)
	} else {
		fmt.Println("Server stopped")
	}
}
