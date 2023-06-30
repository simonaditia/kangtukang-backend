package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/models"
)

type CreateCategoryInput struct {
	Kategori string `json:"kategori" binding:"required"`
}

func GetCategory(c *gin.Context) {
	// db := initDatabase()

	var category models.Category
	if err := models.DB.Preload("Users").First(&category, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(200, category)
}

func AddCategoryUser(c *gin.Context) {
	var category models.Category
	if err := models.DB.First(&category, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	var user models.User
	if err := models.DB.First(&user, c.Param("userID")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	models.DB.Model(&category).Association("Users").Append(&user)

	c.JSON(200, gin.H{"message": "User added to category"})
}

func UpdateUserCategories(c *gin.Context) {
	// db := initDatabase()

	var user models.User
	if err := models.DB.Preload("Categories").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	// Clear existing categories
	models.DB.Model(&user).Association("Categories").Clear()

	// Update categories
	for _, category := range updatedUser.Categories {
		var existingCategory models.Category
		if err := models.DB.First(&existingCategory, category.ID).Error; err != nil {
			c.JSON(404, gin.H{"error": "Category not found"})
			return
		}
		models.DB.Model(&user).Association("Categories").Append(&existingCategory)
	}

	// Fetch the updated user with categories
	if err := models.DB.Preload("Categories").First(&user, user.ID).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, user)
}

/*func CreateCategory(c *gin.Context) {
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
*/
