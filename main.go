package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jordhan-carvalho/aldebaran/internal/db"
)

func main() {

	initDotEnv()
  database.ConnectDB()

	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Aldebaran lives!")
	})

	router.Run()
}

func initDotEnv() {
	fmt.Println("Initializing dot env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
