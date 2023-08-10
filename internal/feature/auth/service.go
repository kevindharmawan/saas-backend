package auth

import (
	"github.com/kevindharmawan/saas-backend/internal/feature/user"
	"github.com/kevindharmawan/saas-backend/internal/model"
)

// TODO: Add verification feature.
// TODO: Check if email is in correct format
type authServiceImpl struct {
	authRepository  AuthRepository
	hashRepository  HashRepository
	tokenRepository TokenRepository
	userRepository  user.UserRepository
}

func NewAuthService(
	authRepository AuthRepository,
	hashRepository HashRepository,
	tokenRepository TokenRepository,
	userRepository user.UserRepository,
) AuthService {
	return &authServiceImpl{
		authRepository:  authRepository,
		hashRepository:  hashRepository,
		tokenRepository: tokenRepository,
		userRepository:  userRepository,
	}
}

func (s *authServiceImpl) SignUp(email string, password string) (*model.Token, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *authServiceImpl) SignIn(email string, password string) (*model.Token, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *authServiceImpl) ValidateToken(token string) (string, *model.AppError) {
	return "", model.NewInternalServerError()
}
