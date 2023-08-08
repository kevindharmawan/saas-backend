package pkg

import (
	"github.com/kevindharmawan/saas-backend/internal/shared/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite(dbConfig config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbConfig.FileName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
