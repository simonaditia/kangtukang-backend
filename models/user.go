package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID         uint   `json:"id" gorm:"primaryKey"`
	Nama      string `json:"nama"`
	NoTelp    string `json:"no_telp"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Alamat    string `json:"alamat" gorm:"type:text"`
	Role      string `json:"role"`
	Entries   []Entry
	Kategori  string  `json:"kategori"`
	Biaya     float64 `json:"biaya"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
	// Categories   []Category `json:"categories" gorm:"many2many:tukang_categories"`
	// CategoriesID []int      `json:"categories_id" form:"categories_id" gorm:"-"`
	// Kategori []string `json:"kategori" gorm:"type:json"`
	// IDKategoriTukang int32  `gorm:"default:null"`
}

func (user *User) Save() (*User, error) {
	err := DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := DB.Preload("Entries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// var user models.User
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Record not found!",
// 		})
// 		return
// 	}
