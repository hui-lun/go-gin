package src

import (
	session "golangAPI/middlewares"
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())

	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUsers)
	//delete user
	//user.DELETE("/:id", service.DeleteUser)
	//put user
	user.PUT("/:id", service.PutUser)
	//login
	user.POST("/login", service.LoginUser)

	// Check User Session
	user.GET("/check", service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		// delete user
		user.DELETE("/:id", service.DeleteUser)
		// Logout user
		user.GET("/logout", service.LogoutUser)
	}
}
