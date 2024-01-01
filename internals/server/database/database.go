package database

import (
	"HRMS/internals/server/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "root"
	dbname   = "hrms"
)

type Database struct {
	dbClient *sql.DB
}

var DB Database

// dbInterface will have custom/new methods for mysql
type dbInterface interface {
}

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
