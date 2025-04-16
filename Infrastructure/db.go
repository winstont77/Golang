// infrastructure/db.go
package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"GolangIrisMVC/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=127.0.0.1 user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{}, &model.Order{}) // 可以加更多 model
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}
}