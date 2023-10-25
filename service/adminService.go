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

type AdminService interface {
	CreateAdmin(srv echo.Context, request modelsrequest.AdminCreate) (*domain.Admin, error)
	LoginAdmin(srv echo.Context, request modelsrequest.AdminLogin) (*domain.Admin, error)
	UpdateAdmin(srv echo.Context, request modelsrequest.AdminUpdate, id int) (*domain.Admin, error)
	DeleteAdmin(srv echo.Context, id int) error
	FindById(srv echo.Context, id int) (*domain.Admin, error)
	FindAll(srv echo.Context) ([]domain.Admin, error)
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminService(AdminRepository repository.AdminRepository, Validate *validator.Validate) *AdminServiceImpl {
	return &AdminServiceImpl{
		AdminRepository: AdminRepository,
		Validate:        Validate,
	}
}

func (service *AdminServiceImpl) CreateAdmin(srv echo.Context, request modelsrequest.AdminCreate) (*domain.Admin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingAdmin, _ := service.AdminRepository.FindByEmail(request.Email)
	if existingAdmin != nil {
		return nil, fmt.Errorf("email Already Exists")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	result, err := service.AdminRepository.Create(admin)
	if err != nil {
		return nil, fmt.Errorf("error when creating admin: %s", err.Error())
	}

	return result, nil

}

func (service *AdminServiceImpl) LoginAdmin(srv echo.Context, request modelsrequest.AdminLogin) (*domain.Admin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingAdmin, err := service.AdminRepository.FindByEmail(request.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	admin := req.AdminLoginRequestToAdminDomain(request)
	err = helper.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid Email or Password")
	}
	return existingAdmin, nil
}

func (service *AdminServiceImpl) UpdateAdmin(srv echo.Context, request modelsrequest.AdminUpdate, id int) (*domain.Admin, error) {

	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	admin.Password = helper.HashPassword(admin.Password)

	result, err := service.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating admin: %s", err.Error())
	}

	return result, nil
}

func (service *AdminServiceImpl) DeleteAdmin(srv echo.Context, id int) error {

	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}

	err := service.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting admin: %s", err)
	}

	return nil
}

func (service *AdminServiceImpl) FindById(srv echo.Context, id int) (*domain.Admin, error) {

	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return existingAdmin, nil
}

func (service *AdminServiceImpl) FindAll(srv echo.Context) ([]domain.Admin, error) {
	admin, err := service.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}

	return admin, nil
}
