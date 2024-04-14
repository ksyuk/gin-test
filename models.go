package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null" size:"255"`
	Items []Item `json:"items"`
}

type Item struct {
	gorm.Model
	Name string `json:"name" gorm:"not null" size:"255"`
	Price int `json:"price" gorm:"not null"`
	UserID uint
}
