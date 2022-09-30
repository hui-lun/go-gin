package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB

var err error

func DD() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:07188@tcp(127.0.0.1:3306)/Demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}
