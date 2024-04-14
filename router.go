package main

import (
	"github.com/gin-gonic/gin"
)

func router () {
	r := gin.Default()

	r.GET("/item/:id", getItemHandler)

    r.POST("/item", createItemHandler)
	r.GET("/user/:id", getUsersHandler)

	r.POST("/user", createUserHandler)

	r.Run()
}
