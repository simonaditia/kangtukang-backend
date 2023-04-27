package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/helper"
	"github.com/simonaditia/kangtukang-backend/models"
)

type CreateUserInput struct {
	Nama     string `json:"nama" binding:"required"`
	NoTelp   string `json:"no_telp" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Alamat   string `json:"alamat"`
	// IDKategoriTukang int32  `json:"id_kategori_tukang"`
}

type UpdateUserInput struct {
	Nama     string `json:"nama"`
	NoTelp   string `json:"no_telp"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
	// IDKategoriTukang int32  `json:"id_kategori_tukang"`
}

// GET /users
// Find all users
func FindUsers(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func Register(context *gin.Context) {
	var input models.AuthenticationInputRegister
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{
		Nama:     input.Nama,
		Email:    input.Email,
		Password: input.Password,
	}

	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"user": savedUser,
	})
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error1": err.Error(),
		})
		return
	}

	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error2": err.Error(),
		})
		return
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error3": err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error4": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"status":  http.StatusOK, //200
		"jwt":     jwt,
	})
}

// POST /users
// Create new user
/*func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create user
	user := models.User{Nama: input.Nama, NoTelp: input.NoTelp, Email: input.Email, Password: input.Password, Alamat: input.Alamat}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}
*/

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   true,
	})
}

/*
func postUser(c *gin.Context) {
	item := User{
		Name:    c.PostForm("name"),
		Address: c.PostForm("address"),
	}

	DB.Create(&item)

	c.JSON(200, gin.H{
		"status": "berhasil ngepost",
		"data":   item,
	})
}
*/
