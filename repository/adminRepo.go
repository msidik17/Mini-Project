package repository

import (
	"Mini-Project/models/domain"
	"Mini-Project/models/schema"
	"Mini-Project/utils/request"
	"Mini-Project/utils/response"
	"fmt"

	"gorm.io/gorm"
)

type AdminRepository interface {
	Create(admin *domain.Admin) (*domain.Admin, error)
	Update(admin *domain.Admin, Id int) (*domain.Admin, error)
	Delete(id int) error
	FindById(id int) (*domain.Admin, error)
	FindByEmail(email string) (*domain.Admin, error)
	FindAll() ([]domain.Admin, error)
}

type AdminRepositoryImpl struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{DB: DB}
}

func (repository *AdminRepositoryImpl) Create(admin *domain.Admin) (*domain.Admin, error) {
	adminDb := request.AdminDomainToAdminSchema(*admin)
	result := repository.DB.Create(&adminDb)
	if result.Error != nil {
		return nil, result.Error
	}
	results := response.AdminSchemaToAdminDomain(adminDb)
	fmt.Println(result)
	return results, nil

}

func (repository *AdminRepositoryImpl) Update(admin *domain.Admin, Id int) (*domain.Admin, error) {
	result := repository.DB.Table("admin").Where("id = ?", Id).Updates(domain.Admin{Name: admin.Name, Email: admin.Email, Password: admin.Password})
	if result.Error != nil {
		return nil, result.Error
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *AdminRepositoryImpl) FindById(id int) (*domain.Admin, error) {
	var admin domain.Admin
	result := repository.DB.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	result := repository.DB.Where("email =?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (repository *AdminRepositoryImpl) FindAll() ([]domain.Admin, error) {
	var admin []domain.Admin

	result := repository.DB.Find(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}
