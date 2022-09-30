package main

import (
	"golangAPI/database"
	"golangAPI/middlewares"
	"golangAPI/pojo"
	. "golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogger() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogger()

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{})
	}

	router.Use(gin.Recovery(), middlewares.Logger())
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	go func() {
		database.DD()
	}()
	router.Run(":8000")
}
