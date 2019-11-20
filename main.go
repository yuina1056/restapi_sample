package main

import (
	"log"
	"os"
	"test_restapi/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.Handler(os.Getenv("PORT"))
}
