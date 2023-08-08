package pkg

import (
	"fmt"

	"github.com/kevindharmawan/saas-backend/internal/shared/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.User, dbConfig.Password, dbConfig.SSLMode),
	), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
