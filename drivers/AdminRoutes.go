package routes

import (
	"Mini-Project/handler"
	"Mini-Project/repository"
	"Mini-Project/service"
	"os"

	"github.com/go-playground/validator"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository, validate)
	adminHandler := handler.NewAdminHandler(adminService)

	adminsGroup := e.Group("admin")

	adminsGroup.POST("", adminHandler.RegisterAdminHandler)
	adminsGroup.POST("/login", adminHandler.LoginAdminHandler)

	adminsGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	adminsGroup.GET("", adminHandler.GetAllAdminHandler)
	adminsGroup.GET("/:id", adminHandler.GetAdminByIdHandler)
	adminsGroup.GET("/:email", adminHandler.GetAdminByEmailHandler)
	adminsGroup.POST("/:id", adminHandler.UpdateAdminHandler)
	adminsGroup.DELETE("/:id", adminHandler.DeleteAdminHandler)
}
