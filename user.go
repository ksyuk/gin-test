package main

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

func getUser (db *gorm.DB, id string) (User, error) {
    var user User
    err := db.First(&user, id).Error
    return user, err
}

func createUser (db *gorm.DB, user User) error {
    err := db.Create(&user).Error
    return err
}

func getUsersHandler(c *gin.Context) {
    id := c.Param("id")

    user, err := getUser(db, id)
    if err != nil {
        c.JSON(500, gin.H{
            "message": err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "user": user,
    })
}

func createUserHandler(c *gin.Context){
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err})
        return
    }

    err := createUser(db, user)
    if err != nil {
        c.JSON(500, gin.H{
            "message": err,
        })
        return
    }

    c.JSON(200, gin.H{
        "message": "User created successfully",
    })
}