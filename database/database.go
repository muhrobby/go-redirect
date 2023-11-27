package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "root:@tcp(127.0.0.1:8111)/go-redirect?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to open database")
	}

	fmt.Println("connecting to database")

	return db

}
