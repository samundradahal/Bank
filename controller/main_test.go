package controller

import (
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func TestMain(m *testing.M) {

	db, er := gorm.Open(mysql.Open("root:root@tcp(localhost:3307)/bank?parseTime=true"), &gorm.Config{})

	if er != nil {
		panic("Error connecting Database")
	}
	DB = db
	os.Exit(m.Run())
}
