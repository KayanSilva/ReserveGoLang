package database

import (
	"github.com/KayanSilva/ReserveGoLang/api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&models.Personality{})
}
