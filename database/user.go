package database

import (
	"echo-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=123456 dbname=intern port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	AutoMigration()

	log.Println("Database Connected.")
}

func AutoMigration() {
	if err := DB.AutoMigrate(
		&models.User{},
		&models.UserAddress{},
		&models.Shop{},
		&models.Product{},
		&models.Category{},
		&models.ProductsCategories{},
		&models.Courier{},
		&models.OrderStatus{},
		&models.Order{},
		&models.ProductOrder{},
	); err != nil {
		panic(err)
	}

	log.Println("AutoAutoMigration completed.")
}
