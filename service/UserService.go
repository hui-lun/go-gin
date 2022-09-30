package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var userList = []pojo.User{}

// Get User
func FindAllUsers(c *gin.Context) {
	//c.JSON(http.StatusOK, userList)
	user := pojo.FindAllUsers()
	c.JSON(http.StatusOK, user)
}

// Get User by Id
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

// Post
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

// Delete
func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// Put
func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
