package handler

import (
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/service"
	"Mini-Project/utils/helper"
	res "Mini-Project/utils/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	CreateOrder(h echo.Context) error
	FindAll(h echo.Context) error
	FindByID(h echo.Context) error
	Delete(h echo.Context) error
}

type OrderHandlerImpl struct {
	OrderService service.OrderService
}

func NewOrderHandler(OrderService service.OrderService) OrderHandler {
	return &OrderHandlerImpl{OrderService: OrderService}
}

func (h *OrderHandlerImpl) CreateOrder(srv echo.Context) error {
	createOrder := modelsrequest.CreateOrder{}
	err := srv.Bind(&createOrder)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.OrderService.CreateOrder(srv, createOrder)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Order Error"))
	}

	response := res.CreateOrderToOrderResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Order Data", response))
}

func (h *OrderHandlerImpl) FindAll(srv echo.Context) error {
	result, err := h.OrderService.FindAll(srv)
	if err != nil {
		if strings.Contains(err.Error(), "orders not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Orders Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find All Orders Data Error"))
	}

	response := res.ConvertOrderResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Find All Order Data", response))
}

func (h *OrderHandlerImpl) FindByID(srv echo.Context) error {
	OrderId := srv.Param("id")
	OrderIdInt, err := strconv.Atoi(OrderId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.OrderService.FindByID(srv, OrderIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "order not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Order Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find order Data Error"))
	}

	response := res.CreateOrderToOrderResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Find Order Data", response))
}

func (h *OrderHandlerImpl) Delete(srv echo.Context) error {
	OrderId := srv.Param("id")
	OrderIdInt, err := strconv.Atoi(OrderId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.OrderService.Delete(srv, OrderIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("movie Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Order Data Error"))
	}

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Order Data", nil))
}
