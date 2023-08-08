package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type validationRepositoryImpl struct{}

func NewValidationRepository() ValidationRepository {
	return &validationRepositoryImpl{}
}

func (r *validationRepositoryImpl) ValidateToken(token string) (int64, *model.AppError) {
	return 0, model.NewInternalServerError()
}
