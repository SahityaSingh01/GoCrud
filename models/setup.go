package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := GetDBConString()
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func GetDBConString() string {
	user := "root"
	password := "root"
	net := "tcp"
	address := "localhost"
	dbname := "Crud"

	connString := fmt.Sprintf("%v:%v@%v(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, net, address, dbname)
	return connString
}
