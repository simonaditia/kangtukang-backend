package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/models"
)

type CreateCategoryInput struct {
	Kategori string `json:"kategori" binding:"required"`
}

func CreateCategory(c *gin.Context) {
	// Validate input
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user
	user := models.Category{Kategori: input.Kategori}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	err := models.DB.Table("categories").Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "success",
		"data":    categories,
	})
}
