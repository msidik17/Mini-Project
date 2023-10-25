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

func UserRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, validate)
	userHandler := handler.NewUserHandler(userService)

	usersGroup := e.Group("user")

	usersGroup.POST("", userHandler.RegisterUserHandler)
	usersGroup.POST("/login", userHandler.LoginUserHandler)

	usersGroup.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	usersGroup.GET("", userHandler.GetAllUserHandler)
	usersGroup.GET("/:id", userHandler.GetUserByIdHandler)
	usersGroup.PUT("/:id", userHandler.UpdateUserHandler)
	usersGroup.DELETE("/:id", userHandler.DeleteUserHandler)
}
