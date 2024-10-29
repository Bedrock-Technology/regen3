package models

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestModelsInit(t *testing.T) {
	// create schema eigen collate utf8mb4_unicode_ci;
	dsn := "root:123456@tcp(127.0.0.1:3306)/eigen?charset=utf8mb4&parseTime=True&loc=UTC"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&Validator{}, &Pod{}, &CheckPoint{},
		&QueueWithdrawals{}, &Cursor{}, &Transaction{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	defer sqlDB.Close()
}
