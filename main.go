package main

import (
	"Mini-Project/configs"
	routes "Mini-Project/drivers"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	validate := validator.New()
	DB := configs.ConnectDB()

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome To Alta Cinema Service")
	})

	routes.AdminRoutes(app, DB, validate)
	routes.UserRoutes(app, DB, validate)

	app.Pre(middleware.RemoveTrailingSlash())
	app.Use(middleware.CORS())
	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	app.Logger.Fatal(app.Start(":8000"))
}
