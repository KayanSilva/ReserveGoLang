package test

import (
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectDatabase(t *testing.T) {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao conectar com o BD: %v", err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		t.Fatalf("Erro ao obter uma conexão com o banco de dados: %v", err)
	}

	if err := sqlDb.Ping(); err != nil {
		t.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	defer func() {
		if err := sqlDb.Close(); err != nil {
			t.Fatalf("Erro ao fechar o banco de dados: %v", err)
		}
	}()
}
