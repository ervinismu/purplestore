package controllers

import (
	"net/http"

	"github.com/ervinismu/purplestore/models"
	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func ListProducts(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{ "data": products })
}

func ShowProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")
	if err := models.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found!" })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "data": product })
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	product := models.Product{
		Name: input.Name,
		Description: input.Description,
	}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{ "data": product })
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")
	if err := models.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found!" })
		return
	}

	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	models.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H {"data": product})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")
	if err := models.DB.Where("id = ?", productId).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found" })
		return
	}

	models.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H { "data": true })
}
