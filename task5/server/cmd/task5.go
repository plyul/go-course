package main

import (
	"github.com/plyul/go-course/task5/server/internal/application"
	"log"
)

func main() {
	app := application.New(":8080")
	log.Printf("application: %v", app.Run())
	log.Println("Exit")
}
