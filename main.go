package main

import (
	"github.com/ervinismu/purplestore/controllers"
	"github.com/ervinismu/purplestore/models"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	models.ConnectDB()

	r.GET("/users", controllers.ListUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.ShowUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.Run()
}
