package main

import (
	"github.com/ervinismu/purplestore/controllers"
	"github.com/ervinismu/purplestore/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	err := godotenv.Load(".env"); if err != nil {
		panic("cannot load env varaiable")
	}

	r := gin.Default()
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

	err = r.Run(); if err != nil {
		panic("Cannot start app")
	}
}
