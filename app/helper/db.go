package helper

import (
	"fmt"
	"log"
	"os"

	// models "app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// connectDb
func ConnectDb() {
	username := "root"
	password := GetEnv("MYSQL_ROOT_PASSWORD", "abc123")
	database := GetEnv("MYSQL_DATABASE", "appdb")
	dbInstance := GetEnv("MYSQL_INSTANCE", "127.0.0.1")
	port := GetEnv("MYSQL_PORT", "3333")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, dbInstance, port, database)
	log.Println("MYSQL:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	DBConn = db

}

// func migrate() {
// 	ConnectDb()
// 	DBConn.AutoMigrate(&models.User{}, &models.Blog{}, &models.Comment{})
// }
