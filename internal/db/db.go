package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"untitled9/internal/config"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	fmt.Println(config.Port, config.Host, config.DatabaseURL, config.User, config.Password)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DatabaseURL,
	)

	fmt.Println("DSN:", dsn)
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("✅ Підключення до PostgreSQL успішне!")
	}

	return DB
}
