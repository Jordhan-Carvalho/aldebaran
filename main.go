package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/jordhan-carvalho/aldebaran/internal/db"
	"github.com/jordhan-carvalho/aldebaran/internal"
)

func main() {

	initDotEnv()
  database.ConnectDB()
  webserver.InitWebServer()

}

func initDotEnv() {
	fmt.Println("Initializing dot env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
