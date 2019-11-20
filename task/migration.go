package main

import (
	"log"
	"test_restapi/db"
)

func main() {
	db.Open()
	defer db.Close()
	db.AutoMigrate()
	log.Printf("Migration Completed!")
}
