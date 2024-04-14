package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
}

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

	return DB
}
