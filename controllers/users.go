package controllers

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/simonaditia/kangtukang-backend/helper"
	"github.com/simonaditia/kangtukang-backend/models"
	"googlemaps.github.io/maps"
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
	Nama      string  `json:"nama"`
	NoTelp    string  `json:"no_telp"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Alamat    string  `json:"alamat"`
	Kategori  string  `json:"kategori"`
	Biaya     float64 `json:"biaya"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	ImageUrl  string  `json:"image_url"`
	// CategoriesID []int  `json:"categories_id"`
	// CategoriesID []int `json:"categories_id" form:"categories_id" gorm:"-"`
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
	if err := models.DB.Preload("Categories").First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	// address, err := helper.TranslateToAddress(user.Latitude, user.Longitude)
	// if err != nil {
	// 	panic(err)
	// }

	// user.Alamat = address

	// if err := models.DB.Save(&user).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "Failed to save changes!",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func AddUserCategory(c *gin.Context) {
	// db := initDatabase()

	var user models.User
	if err := models.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var category models.Category
	if err := models.DB.First(&category, c.Param("categoryID")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}

	models.DB.Model(&user).Association("Categories").Append(&category)

	c.JSON(200, gin.H{"message": "Category added to user"})
}

func FindUserByEmail(c *gin.Context) {
	var users []models.User
	var email string = c.Query("email")

	if email == "" {
		err := models.DB.Table("users").Find(&users).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	err := models.DB.Table("users").
		Where("email LIKE ?", "%"+email+"%").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}

func CheckIsAvailableEmail(c *gin.Context) {
	var users []models.User
	var email string = c.Query("email")

	if email == "" {
		err := models.DB.Table("users").Find(&users).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	} else {
		// Lakukan validasi jika email sudah ada dalam basis data
		// Misalnya, menggunakan ORM seperti Gorm atau query langsung ke basis data

		// Contoh validasi dengan menggunakan Gorm
		var count int64
		err := models.DB.Table("users").Where("email = ?", email).Count(&count).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":     "Email already exists",
				"available": false,
			})
			return
		}

		err = models.DB.Table("users").
			Where("email LIKE ?", "%"+email+"%").Find(&users).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"data":      users,
		"available": true,
	})
}

/*func FindTukang(c *gin.Context) {
	var user models.User
	var nama string = c.Query("nama")
	fmt.Println("halo findtukang")
	var kategori string = c.Query("kategori")
	fmt.Println(nama, kategori)
	// select user.*, kategori_tukang.kategori from user inner join kategori_tukang on kategori_tukang.tukang_id = user.id where user.role = 'tukang' AND user.nama like 'adit';

	// select users.*, tukang_categories.kategori from users inner join tukang_categories on tukang_categories.tukang_id = users.id where users.role = 'tukang' AND (users.nama like 'james-tukang' OR tukang_categories.kategori like 'renovasi');

	err := models.DB.Table("users").
		Select("users.*, tukang_categories.kategori").
		Joins("INNER JOIN tukang_categories ON tukang_categories.tukang_id = users.id").
		Where("users.role LIKE ?", "%tukang%").
		Where("users.nama LIKE ? OR tukang_categories.kategori LIKE ?", "%"+nama+"%", "%"+kategori+"%").Find(&user).Error
	// err := models.DB.Table("users").Select("users.id").Where("users.role = 'tukang'").Scan(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// if err := models.DB.Where("role = ?", "tukang").Find(&user).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Record not found!",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}*/

/*func FindTukang(c *gin.Context) {
	var users []models.User
	var nama string = c.Query("nama")
	var kategori string = c.Query("kategori")
	fmt.Println(nama, kategori)

	err := models.DB.Table("users").
		// Select("users.*, tukang_categories.kategori").
		// select users.*, GROUP_CONCAT(tukang_categories.kategori SEPARATOR ',') as kategori from users inner join tukang_categories on tukang_categories.tukang_id = users.id where users.role = 'tukang' AND (users.nama like '%james-tukang2%' OR tukang_categories.kategori like 'renovasi' GROUP(users.id));

		Select("users.*, GROUP_CONCAT(tukang_categories.kategori SEPARATOR ',') as kategori").
		Joins("INNER JOIN tukang_categories ON tukang_categories.tukang_id = users.id").
		Where("users.role LIKE ?", "%tukang%").
		Where("users.nama LIKE ? OR tukang_categories.kategori LIKE ?", "%"+nama+"%", "%"+kategori+"%").
		Group("users.id").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   users,
	})
}*/

func FindTukang(c *gin.Context) {
	var users []models.User
	var nama string = c.Query("nama")
	var kategori string = c.Query("kategori")
	const ROLE = "tukang"
	// fmt.Println(nama, kategori)

	// 1. Parse Token JWT dan Ambil Nilai Latitude dan Longitude User Customer
	latitude, longitude, err := helper.ParseLatitudeLongitudeFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse latitude and longitude from token",
		})
		return
	}
	fmt.Println(latitude, longitude)

	if nama == "" {
		err := models.DB.Table("users").Find(&users).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	// err = models.DB.Table("users").
	// 	Where("nama LIKE ? AND kategori LIKE ? AND role = ?", "%"+nama+"%", "%"+kategori+"%", ROLE).Find(&users).Error

	err = models.DB.Table("users").
		Select("DISTINCT users.*").
		Joins("JOIN user_categories ON user_categories.user_id = users.id").
		Joins("JOIN categories ON categories.id = user_categories.category_id").
		Where("users.nama LIKE ? AND categories.name LIKE ? AND users.role = ?", "%"+nama+"%", "%"+kategori+"%", ROLE).
		Preload("Categories").
		Find(&users).
		Error

	var tukangs []models.User
	// var distance float64
	for _, user := range users {
		// fmt.Println("Tampil user", user)
		start := &maps.LatLng{Lat: latitude, Lng: longitude}
		end := &maps.LatLng{Lat: user.Latitude, Lng: user.Longitude}

		result, err := models.FindTukangWithAStarGmaps(start, end)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		user.Distance = result
		tukangs = append(tukangs, user)
	}

	sort.Slice(tukangs, func(i, j int) bool {
		return tukangs[i].Distance < tukangs[j].Distance
	})

	/*err := models.DB.Preload("").Find(&users).Error
	result := make([]gin.H, len(users))*/
	// for k, kategori := range users {
	// 	kategories := make([]string, len(kategori))
	// }
	/*for i, user := range users {
		categories := make([]string, len(user.Categories))
		for j, category := range user.Categories {
			categories[j] = category.Kategori
		}
		fmt.Println(categories)
		result[i] = gin.H{
			"Nama":     user.Nama,
			"Email":    user.Email,
			"Kategori": categories,
		}
	}*/

	// Find(&users, "nama = ?", nama, "role = ?", ROLE).Error
	// Select("users.*, tukang_categories.kategori").
	// Joins("INNER JOIN tukang_categories ON tukang_categories.tukang_id = users.id").
	// // Where("users.role LIKE ?", "%tukang%").
	// Where("users.nama LIKE ? OR tukang_categories.kategori LIKE ?", "%"+nama+"%", "%"+kategori+"%").
	// Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tukangs,
		// "data":   users,
	})
}

func DetailTukang(c *gin.Context) {
	var user models.User
	const ROLE = "tukang"
	// err := models.DB.Table("users").
	// 	Where("nama LIKE ? AND kategori LIKE ? AND role = ?", "%"+nama+"%", "%"+kategori+"%", ROLE).Find(&users).Error
	if err := models.DB.Where("id = ? AND role = ?", c.Param("id"), ROLE).First(&user).Error; err != nil {
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

func RegisterCustomer(context *gin.Context) {
	var input models.AuthenticationInputRegister
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	address, err := helper.TranslateToAddress(input.Latitude, input.Longitude)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Nama:      input.Nama,
		Email:     input.Email,
		Password:  input.Password,
		Role:      "customer",
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Alamat:    address,
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

func RegisterTukang(context *gin.Context) {
	var input models.AuthenticationInputRegister
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	address, err := helper.TranslateToAddress(input.Latitude, input.Longitude)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Nama:      input.Nama,
		Email:     input.Email,
		Password:  input.Password,
		Role:      "tukang",
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		// Kategori:  "Renovasi",
		Alamat: address,
		Biaya:  100000,
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
			"error": err.Error(),
		})
		return
	}

	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(jwt)

	context.JSON(http.StatusOK, gin.H{
		"user":    user,
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
	// var tukangCategor models.TukangCategory
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
	// fmt.Println("user.CategoriesID", user.CategoriesID)
	// fmt.Println("input.CategoriesID", input.CategoriesID)

	/*for _, categoryID := range input.CategoriesID {
		// userCategory := new(models.TukangCategory)
		tukangCategory.UserID = user.ID
		tukangCategory.CategoryID = categoryID
		// models.DB.Model(&tukangCategory).Updates(input)
		// models.DB.Delete(&tukangCategory)
		// models.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&tukangCategory)
		models.DB.Create(&tukangCategory)
		models.DB.Model(&tukangCategory).Updates(tukangCategory.UserID)
		models.DB.Model(&tukangCategory).Updates(tukangCategory.CategoryID)
		// models.DB.Model(&tukangCategory).Updates(tukangCategory).Where()
	}*/

	/*for _, categoryID := range user.CategoriesID {
		userCategory := new(models.TukangCategory)
		userCategory.UserID = user.ID
		userCategory.CategoryID = categoryID
		// models.DB.Create(&userCategory)
		models.DB.Model(&userCategory).Updates(input)
	}*/

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
		// "data2":  input.CategoriesID,
		// "id":            user.ID,
		// "nama":          user.Nama,
		// "email":         user.Email,
		// "no_telp":       user.NoTelp,
		// "alamat":        user.Alamat,
		// "categories_id": input.CategoriesID,
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
