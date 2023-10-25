package handler

import (
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/service"
	"Mini-Project/utils/helper"
	res "Mini-Project/utils/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	RegisterAdminHandler(srv echo.Context) error
	LoginAdminHandler(srv echo.Context) error
	UpdateAdminHandler(srv echo.Context) error
	DeleteAdminHandler(srv echo.Context) error
	GetAllAdminHandler(srv echo.Context) error
	GetAdminByIdHandler(srv echo.Context) error
}

type AdminHandlerImpl struct {
	AdminService service.AdminService
}

func NewAdminHandler(AdminService service.AdminService) AdminHandler {
	return &AdminHandlerImpl{AdminService: AdminService}
}

func (h *AdminHandlerImpl) RegisterAdminHandler(srv echo.Context) error {
	adminCreateRequest := modelsrequest.AdminCreate{}
	err := srv.Bind(&adminCreateRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.AdminService.CreateAdmin(srv, adminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "email already exists") {
			return srv.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exists"))
		}
		fmt.Println(result)
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}
	response := res.AdminDomaintoAdminResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (h *AdminHandlerImpl) LoginAdminHandler(srv echo.Context) error {
	adminLoginRequest := modelsrequest.AdminLogin{}
	err := srv.Bind(&adminLoginRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := h.AdminService.LoginAdmin(srv, adminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}
		if strings.Contains(err.Error(), "email already exists") {
			return srv.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exists"))
		}
		fmt.Println(response)
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}
	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := helper.GenerateAdminToken(&adminLoginResponse, uint(response.ID))
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	adminLoginResponse.Token = token

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

func (h *AdminHandlerImpl) UpdateAdminHandler(srv echo.Context) error {
	adminId := srv.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	adminUpdateRequest := modelsrequest.AdminUpdate{}
	err = srv.Bind(&adminUpdateRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.AdminService.UpdateAdmin(srv, adminUpdateRequest, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "admin not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)
	fmt.Print(result)
	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Admin Data", response))
}

func (h *AdminHandlerImpl) DeleteAdminHandler(srv echo.Context) error {
	adminId := srv.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.AdminService.DeleteAdmin(srv, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Admin Data", nil))
}

func (h *AdminHandlerImpl) GetAllAdminHandler(srv echo.Context) error {
	result, err := h.AdminService.FindAll(srv)
	if err != nil {
		if strings.Contains(err.Error(), "admins not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Admins Data Error"))
	}

	response := res.ConvertAdminResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get All Admin Data", response))
}

func (h *AdminHandlerImpl) GetAdminByIdHandler(srv echo.Context) error {
	adminId := srv.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.AdminService.FindById(srv, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", response))
}