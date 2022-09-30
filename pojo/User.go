package pojo

import (
	"golangAPI/database"
	"log"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DBconnect.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	database.DBconnect.Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user User) User {
	database.DBconnect.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBconnect.Where("id = ?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId string, user User) User {
	database.DBconnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
