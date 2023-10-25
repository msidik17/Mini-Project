package service

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/repository"
	"Mini-Project/utils/helper"
	req "Mini-Project/utils/request"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(srv echo.Context, request modelsrequest.UserCreate) (*domain.User, error)
	LoginUser(srv echo.Context, request modelsrequest.UserLogin) (*domain.User, error)
	UpdateUser(srv echo.Context, request modelsrequest.UserUpdate, id int) (*domain.User, error)
	Deleteuser(srv echo.Context, id int) error
	FindById(srv echo.Context, id int) (*domain.User, error)
	FindAll(srv echo.Context) ([]domain.User, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(UserRepository repository.UserRepository, Validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: UserRepository,
		Validate:        Validate,
	}
}

func (service *UserServiceImpl) CreateUser(srv echo.Context, request modelsrequest.UserCreate) (*domain.User, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingUser, _ := service.UserRepository.FindByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("email Already Exists")
	}

	user := req.UserCreateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error when creating user: %s", err.Error())
	}

	return result, nil

}

func (service *UserServiceImpl) LoginUser(srv echo.Context, request modelsrequest.UserLogin) (*domain.User, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingUser, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	user := req.UserLoginRequestToUserDomain(request)
	err = helper.ComparePassword(existingUser.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid Email or Password")
	}
	return existingUser, nil
}

func (service *UserServiceImpl) UpdateUser(srv echo.Context, request modelsrequest.UserUpdate, id int) (*domain.User, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	user := req.UserUpdateRequestToUserDomain(request)
	user.Password = helper.HashPassword(user.Password)

	result, err := service.UserRepository.Update(user, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating user: %s", err.Error())
	}

	return result, nil
}

func (service *UserServiceImpl) Deleteuser(srv echo.Context, id int) error {

	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return fmt.Errorf("user not found")
	}

	err := service.UserRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting user: %s", err)
	}

	return nil
}

func (service *UserServiceImpl) FindById(srv echo.Context, id int) (*domain.User, error) {

	existingUser, _ := service.UserRepository.FindById(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	return existingUser, nil
}

func (service *UserServiceImpl) FindAll(srv echo.Context) ([]domain.User, error) {
	user, err := service.UserRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("users not found")
	}

	return user, nil
}
