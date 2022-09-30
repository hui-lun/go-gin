package main

import (
	"golangAPI/database"
	"golangAPI/middlewares"
	. "golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogger() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogger()

	router := gin.Default()
	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger())
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DD()
	}()
	router.Run(":8000")
}
