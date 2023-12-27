package main

import (
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

func CreateCon() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
	DB = db
	return DB
}
func main() {
	CreateCon()
}
