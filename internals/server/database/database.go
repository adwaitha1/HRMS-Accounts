package database

import (
	"HRMS/internals/server/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "root"
	dbname   = "hrms"
)

func CreateCon(dbConfig config.DatabaseConf) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Error opening database connection: %v", err)
		return nil, err
	}
	DB = db
	// Check if the connection is successful.
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}
	fmt.Println("Connected to the database!")
	// Enable query logging
	return db, nil
}
