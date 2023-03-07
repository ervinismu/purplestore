package main

import (
	"github.com/ervinismu/purplestore/controllers"
	"github.com/ervinismu/purplestore/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8003", "http://0.0.0.0:8003", "http://127.0.0.1:8003"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
	  }))

	models.ConnectDB()

	r.GET("/users", controllers.ListUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.ShowUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.GET("/products", controllers.ListProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.ShowProduct)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	err := r.Run(); if err != nil {
		panic("Cannot start app")
	}
}
