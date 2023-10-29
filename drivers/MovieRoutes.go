package routes

import (
	"Mini-Project/handler"
	"Mini-Project/repository"
	"Mini-Project/service"
	"Mini-Project/utils/helper"
	"os"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MovieRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository, validate)
	movieHandler := handler.NewMovieHandler(movieService)

	moviesGroup := e.Group("")
	// admin := e.Group("", helper.VerifySignIn("ADMIN_JWT_SECRET"))
	moviesGroup.Use(echoJwt.JWT([]byte(os.Getenv("JWT_SECRET"))))

	//Movies
	moviesGroup.GET("/movie", movieHandler.FindAll)
	moviesGroup.GET("/movie/:id", movieHandler.FindMovieByID)
	moviesGroup.GET("/movie/title/:title", movieHandler.FindMovieByTitle)
	moviesGroup.POST("/admin/movie", movieHandler.AddMovie, helper.AuthMiddleware("admin"))
	moviesGroup.PUT("/admin/movie/:id", movieHandler.UpdateMovie, helper.AuthMiddleware("admin"))
	moviesGroup.DELETE("/admin/movie/:id", movieHandler.DeleteMovie, helper.AuthMiddleware("admin"))

}
