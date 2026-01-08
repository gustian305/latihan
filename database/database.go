package database

import (
	"fmt"
	"latihan/config"
	"latihan/models"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

// Koneksi untuk MySQL
func ConnectMySQL(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke MySQL:", err)
	}

	fmt.Println("Berhasil terkoneksi ke MySQL")
	
	// Auto Migrate
	db.AutoMigrate(&models.Student{})
	DB = db
}

// Koneksi untuk PostgreSQL
func ConnectPostgres(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke PostgreSQL:", err)
	}

	fmt.Println("Berhasil terkoneksi ke PostgreSQL")

	// Auto Migrate
	db.AutoMigrate(&models.Student{})
	DB = db
}