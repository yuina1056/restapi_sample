package main

import (
	"os"
	"test_restapi/server"
)

func main() {
	server.Handler(os.Getenv("PORT"))
}
