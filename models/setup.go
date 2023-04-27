package models

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	fmt.Println(host, username, password, databaseName, port)
	// dsn := "root:@tcp(127.0.0.1:3306)/kangtukang_db?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, host, port, databaseName)
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, username, password, databaseName, port)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	// mysqlDB, err := DB.DB()
	// if err != nil {
	// 	panic(err)
	// }
	// defer mysqlDB.Close()
	return DB
}
