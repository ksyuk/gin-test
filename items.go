package main

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func getItem (db *gorm.DB, id string) (Item, error) {
    var item Item
    err := db.First(&item, id).Error
    return item, err
}

func createItem (db *gorm.DB, item Item) error {
    err := db.Create(&item).Error
    return err
}

func getItemHandler(c *gin.Context) {
	id := c.Param("id")

	item, err := getItem(db, id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"item": item,
	})
}
func createItemHandler (c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	err := createItem(db, item)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Item created successfully",
	})
}
