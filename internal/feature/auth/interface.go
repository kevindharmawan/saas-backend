package auth

import "github.com/kevindharmawan/saas-backend/internal/model"

type AuthRepository interface {
	CreateAuthUser(*model.Auth) (int64, *model.AppError)
	GetAuthUserByEmail(email string) (*model.Auth, *model.AppError)
}

type HashRepository interface {
	HashString(str string) (string, *model.AppError)
	MatchHashAndString(hashedStr string, str string) (bool, *model.AppError)
}

type TokenRepository interface {
	CreateToken(authId int64) (*model.Token, *model.AppError)
}

type ValidationRepository interface {
	ValidateToken(token string) (int64, *model.AppError)
}

type ValidationService interface {
	ValidateToken(token string) (string, *model.AppError)
}

type AuthService interface {
	ValidationService
	SignUp(email string, password string) (*model.Token, *model.AppError)
	SignIn(email string, password string) (*model.Token, *model.AppError)
}
