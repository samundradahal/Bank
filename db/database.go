package db

import (
	"bank/util"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}
	var er error
	db, er := gorm.Open(mysql.Open(config.Dbdriver), &gorm.Config{})

	if er != nil {
		panic("Error connecting Database")
	}
	DB = db
	fmt.Println("DataBase Connected")

}
