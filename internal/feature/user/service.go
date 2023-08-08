package user

import "github.com/kevindharmawan/saas-backend/internal/model"

type userServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) CreateUser(user *model.User) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *userServiceImpl) UpdateUser(user *model.User) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *userServiceImpl) DeleteUser(id int64) *model.AppError {
	return model.NewInternalServerError()
}

func (s *userServiceImpl) GetUserById(id int64) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *userServiceImpl) GetUserByEmail(email string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *userServiceImpl) GetUserByAuthId(authId string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}

func (s *userServiceImpl) GetUserByCustomerId(customerId string) (*model.User, *model.AppError) {
	return nil, model.NewInternalServerError()
}
