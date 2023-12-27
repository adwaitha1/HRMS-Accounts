package main

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabaseConnection(t *testing.T) {

	db := CreateCon()
	if db == nil {
		t.Fatal("Database connection is nil.")
	}
	err := db.Ping()
	if err != nil {
		t.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Database connection test passed!")
}
