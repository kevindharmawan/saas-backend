package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type authRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &authRepositoryImpl{}
}

func (r *authRepositoryImpl) CreateAuthUser(*model.Auth) (int64, *model.AppError) {
	return 0, model.NewInternalServerError()
}

func (r *authRepositoryImpl) GetAuthUserByEmail(email string) (*model.Auth, *model.AppError) {
	return nil, model.NewInternalServerError()
}
