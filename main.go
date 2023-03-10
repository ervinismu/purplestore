package main

import (
	"github.com/ervinismu/purplestore/db"
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

	db := db.ConnectDB()

	userController := InitUserController(db)
	r.GET("/users", userController.ListAllUser)
	r.GET("/users/:id", userController.ShowUser)
	r.POST("/users", userController.CreateUser)
	r.PATCH("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)

	productController := InitProductController(db)
	r.GET("/products", productController.ListProducts)
	r.GET("/products/:id", productController.ShowProduct)
	r.POST("/products", productController.CreateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
	r.PATCH("/products/:id", productController.UpdateProduct)

	err := r.Run(); if err != nil {
		panic("Cannot start app")
	}
}
