package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type tokenRepositoryImpl struct{}

func NewTokenRepository() TokenRepository {
	return &tokenRepositoryImpl{}
}

func (r *tokenRepositoryImpl) CreateToken(authId int64) (*model.Token, *model.AppError) {
	return nil, model.NewInternalServerError()
}
