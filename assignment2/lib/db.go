package lib

import (
	"fmt"
	"os"

	"github.com/aulianurhady/training/assignment2/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() *gorm.DB {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	return db
}

func InitDatabase() {
	db := connectDB()
	migrations.AutoMigrate(db)
}
