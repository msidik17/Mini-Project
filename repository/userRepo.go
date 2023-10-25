package repository

import (
	"Mini-Project/models/domain"
	"Mini-Project/models/schema"
	"Mini-Project/utils/request"
	"Mini-Project/utils/response"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User, Id int) (*domain.User, error)
	Delete(id int) error
	FindById(id int) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll() ([]domain.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) Create(user *domain.User) (*domain.User, error) {
	userDb := request.UserDomainToUserSchema(*user)
	result := repository.DB.Create(&userDb)
	if result.Error != nil {
		return nil, result.Error
	}
	results := response.UserSchemaToUserDomain(userDb)
	fmt.Println(result)
	return results, nil

}

func (repository *UserRepositoryImpl) Update(user *domain.User, id int) (*domain.User, error) {
	result := repository.DB.Table("users").Where("id = ?", id).Updates(domain.User{Name: user.Name, Email: user.Email, Password: user.Password})
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Admin{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *UserRepositoryImpl) FindById(id int) (*domain.User, error) {
	var user domain.User
	result := repository.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}


func (repository *UserRepositoryImpl) FindAll() ([]domain.User, error) {
	var user []domain.User

	result := repository.DB.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
