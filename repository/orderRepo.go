package repository

import (
	"Mini-Project/models/domain"
	"Mini-Project/models/schema"
	"Mini-Project/utils/request"
	"Mini-Project/utils/response"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
	Delete(id int) error
	FindAll() ([]domain.Order, error)
	FindByID(id int) (*domain.Order, error)
}

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{DB: DB}
}

func (repository *OrderRepositoryImpl) CreateOrder(order *domain.Order) (*domain.Order, error) {
	orderDb := request.OrderDomaintoOrderSchema(*order)
	result := repository.DB.Create(&orderDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := response.OrderSchematoOrderDomain(orderDb)

	return results, nil
}


func (repository *OrderRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Order{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *OrderRepositoryImpl) FindAll() ([]domain.Order, error) {
	var order []domain.Order

	if err := repository.DB.Where("deleted_at IS NULL").Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (repository *OrderRepositoryImpl) FindByID(id int) (*domain.Order, error) {
	var order domain.Order
	if err := repository.DB.Unscoped().Where("id = ? AND deleted_at IS NULL", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

