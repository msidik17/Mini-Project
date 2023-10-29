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

type UserHandler interface {
	RegisterUserHandler(h echo.Context) error
	LoginUserHandler(h echo.Context) error
	UpdateUserHandler(h echo.Context) error
	DeleteUserHandler(h echo.Context) error
	GetAllUserHandler( echo.Context) error
	GetUserByIdHandler(h echo.Context) error
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(UserService service.UserService) UserHandler {
	return &UserHandlerImpl{UserService: UserService}
}

func (h *UserHandlerImpl) RegisterUserHandler(srv echo.Context) error {
	userCreateRequest := modelsrequest.UserCreate{}
	err := srv.Bind(&userCreateRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.UserService.CreateUser(srv, userCreateRequest)
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
	response := res.UserDomaintoUserResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (h *UserHandlerImpl) LoginUserHandler(srv echo.Context) error {
	adminLoginRequest := modelsrequest.UserLogin{}
	err := srv.Bind(&adminLoginRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := h.UserService.LoginUser(srv, adminLoginRequest)
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
	userLoginResponse := res.UserDomainToUserLoginResponse(response)

	token, err := helper.GenerateUserToken(&userLoginResponse, uint(response.ID))
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	userLoginResponse.Token = token

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", userLoginResponse))
}

func (h *UserHandlerImpl) UpdateUserHandler(srv echo.Context) error {
	userId := srv.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	userUpdateRequest := modelsrequest.UserUpdate{}
	err = srv.Bind(&userUpdateRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.UserService.UpdateUser(srv, userUpdateRequest, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "user not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Error"))
	}

	response := res.UserDomaintoUserResponse(result)
	fmt.Print(result)
	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated User Data", response))
}

func (h *UserHandlerImpl) DeleteUserHandler(srv echo.Context) error {
	userId := srv.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.UserService.Deleteuser(srv, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete User Data Error"))
	}

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted User Data", nil))
}

func (h *UserHandlerImpl) GetAllUserHandler(srv echo.Context) error {
	result, err := h.UserService.FindAll(srv)
	if err != nil {
		if strings.Contains(err.Error(), "users not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Users Data Error"))
	}

	response := res.ConvertUserResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get All User Data", response))
}

func (h *UserHandlerImpl) GetUserByIdHandler(srv echo.Context) error {
	adminId := srv.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.UserService.FindById(srv, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get User Data", response))
}
