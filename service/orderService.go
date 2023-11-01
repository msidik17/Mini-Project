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

type OrderService interface {
	CreateOrder(srv echo.Context, request modelsrequest.CreateOrder) (*domain.Order, error)
	FindAll(srv echo.Context) ([]domain.Order, error)
	FindByID(srv echo.Context, id int) (*domain.Order, error)
	Delete(srv echo.Context, id int) error
}

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	Validate         *validator.Validate
}

func NewOrderService(OrderRepository repository.OrderRepository, Validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: OrderRepository,
		Validate:         Validate,
	}
}

func (service *OrderServiceImpl) CreateOrder(srv echo.Context, request modelsrequest.CreateOrder) (*domain.Order, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	order := req.OrderCreateRequestToOrderDomain(&request)

	result, err := service.OrderRepository.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("error create order: %s", err.Error())
	}

	return result, nil
}

func (service *OrderServiceImpl) FindAll(srv echo.Context) ([]domain.Order, error) {
	order, err := service.OrderRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("orders not found")
	}

	return order, nil
}

func (service *OrderServiceImpl) FindByID(srv echo.Context, id int) (*domain.Order, error) {
	order, err := service.OrderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (service *OrderServiceImpl) Delete(srv echo.Context, id int) error {

	existingMovie, _ := service.OrderRepository.FindByID(id)
	if existingMovie == nil {
		return fmt.Errorf("order not found")
	}

	err := service.OrderRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting order: %s", err)
	}

	return nil
}
