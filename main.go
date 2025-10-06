package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tekluabayney/taskmanger/app"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := app.New()
	app.Start()
}
