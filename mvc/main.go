package main

import (
	"log"

	"github.com/sharmari/mvc/app"
)

func main() {
	app.StartApp()
	log.Printf("Server Running localhost:8080")
}
