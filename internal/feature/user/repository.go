package user

import (
	"github.com/kevindharmawan/saas-backend/internal/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
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
