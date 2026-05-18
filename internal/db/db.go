package db

import (
	"coffee-app-server/internal/order"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=coffee port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {

		return nil, err
	}

	log.Println("Successfully connected to Postgres!")

	err = db.AutoMigrate(&order.Order{})
	if err != nil {
		return nil, err
	}

	log.Println("Database migration completed successfully!")

	return db, nil
}
