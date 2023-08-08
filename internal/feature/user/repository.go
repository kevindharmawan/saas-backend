package user

import "github.com/kevindharmawan/saas-backend/internal/model"

type userRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) CreateUser(user *model.User) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *userRepositoryImpl) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *userRepositoryImpl) DeleteUser(id int64) *model.AppError {
	return model.NewInternalServerError()
}

func (r *userRepositoryImpl) GetUserById(id int64) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *userRepositoryImpl) GetUserByAuthId(authId string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (r *userRepositoryImpl) GetUserByCustomerId(customerId string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}
