package config

import (
	"log"
	"myapi/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Conexão com o Postgres (usando host "db" pois o docker-compose cria essa rede)
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o BD: %v", err)
	}

	// AutoMigrate para criar/ajustar tabelas
	DB.AutoMigrate(&models.Item{})
	DB.AutoMigrate(&models.Categoria{})
}
