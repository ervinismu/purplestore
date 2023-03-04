package controllers

import (
	"net/http"

	"github.com/ervinismu/purplestore/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Email  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
}

func ListUsers(c *gin.Context) {
	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{ "data": users })
}

func ShowUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found!" })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "data": user })
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	user := models.User{
		Email: input.Email,
		Username: input.Username,
		Password: input.Password,
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{ "data": user })
}

func UpdateUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found!" })
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": err.Error() })
		return
	}

	user.Username = input.Username
	models.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H {"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	userId := c.Param("id")
	if err := models.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H { "error": "Record not found" })
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H { "data": true })
}
