package auth

import (
	"github.com/kevindharmawan/saas-backend/internal/model"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{
		db: db,
	}
}

func (r *authRepositoryImpl) CreateAuthUser(*model.Auth) (int64, *model.AppError) {
	return 0, model.NewInternalServerError()
}

func (r *authRepositoryImpl) GetAuthUserByEmail(email string) (*model.Auth, *model.AppError) {
	return nil, model.NewInternalServerError()
}
