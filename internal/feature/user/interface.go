package user

import "github.com/kevindharmawan/saas-backend/internal/model"

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, *model.AppError)
	UpdateUser(user *model.User) (*model.User, *model.AppError)
	DeleteUser(id int64) *model.AppError
	GetUserById(id int64) (*model.User, *model.AppError)
	GetUserByEmail(email string) (*model.User, *model.AppError)
	GetUserByAuthId(authId string) (*model.User, *model.AppError)
	GetUserByCustomerId(customerId string) (*model.User, *model.AppError)
}

type UserService interface {
	CreateUser(user *model.User) (*model.User, *model.AppError)
	UpdateUser(user *model.User) (*model.User, *model.AppError)
	DeleteUser(id int64) *model.AppError
	GetUserById(id int64) (*model.User, *model.AppError)
	GetUserByEmail(email string) (*model.User, *model.AppError)
	GetUserByAuthId(authId string) (*model.User, *model.AppError)
	GetUserByCustomerId(customerId string) (*model.User, *model.AppError)
}
