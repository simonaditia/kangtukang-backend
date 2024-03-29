package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/simonaditia/kangtukang-backend/models"
	"github.com/simonaditia/kangtukang-backend/routes"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	loadEnv()
	loadDatabase()

	router := routes.SetupRoutes(DB)
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173"}
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	// router.Use(cors.New(config))
	// router.Run()
	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func loadDatabase() {
	DB = models.ConnectDatabase()
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Entry{})
	DB.AutoMigrate(&models.Orders{})
	DB.AutoMigrate(&models.Category{})
	// DB.AutoMigrate(&models.TukangCategory{})
}
