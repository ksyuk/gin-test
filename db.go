package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	var DB *gorm.DB

	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&interpolateParams=true&parseTime=True&loc=Local",
		"localuser", "localpass", "localhost", "localdb")
		// "localuser", "localpass", "db", "localdb")

	err := error(nil)
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func migrate () {
	err := db.AutoMigrate(&User{}, &Item{})
	if err != nil {
		panic(err)
	}
}
