package database

import (
	"fmt"
	"log"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBInstance() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbConfig.Host, dbConfig.Username, dbConfig.Password, dbConfig.DbName, dbConfig.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to make database connection...")
	}
	return db
}
