package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func router () {
	r := gin.Default()

	db := ConnectDB()

	r.GET("/item/:id", func(c *gin.Context) {
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
    })

    r.POST("/item", func(c *gin.Context) {
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
    })

	r.Run()
}

func migrate () {
	db := ConnectDB()
	db.AutoMigrate(&Item{})
}

func main() {
	if len(os.Args) < 2 {
        fmt.Println("No function specified")
        return
    }

    switch os.Args[1] {
    case "router":
        router()
    case "migrate":
        migrate()
    default:
        fmt.Println("Unknown function: " + os.Args[1])
	}
}