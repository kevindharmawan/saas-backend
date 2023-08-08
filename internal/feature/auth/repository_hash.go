package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type hashRepositoryImpl struct{}

func NewHashRepository() HashRepository {
	return &hashRepositoryImpl{}
}

func (r *hashRepositoryImpl) HashString(str string) (string, *model.AppError) {
	return "", model.NewInternalServerError()
}

func (r *hashRepositoryImpl) MatchHashAndString(hashedStr string, str string) (bool, *model.AppError) {
	return false, model.NewInternalServerError()
}
