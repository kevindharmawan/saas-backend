package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type tokenRepositoryImpl struct{}

func NewTokenRepository() TokenRepository {
	return &tokenRepositoryImpl{}
}

func (r *tokenRepositoryImpl) CreateToken(authId int64) (*model.Token, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *tokenRepositoryImpl) ValidateToken(token string) (int64, *model.AppError) {
	return 0, model.NewInternalServerError()
}
