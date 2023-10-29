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

type MovieHandler interface {
	AddMovie(h echo.Context) error
	FindAll(h echo.Context) error
	FindMovieByID(h echo.Context) error
	FindMovieByTitle(h echo.Context) error
	UpdateMovie(h echo.Context) error
	DeleteMovie(h echo.Context) error
}

type MovieHandlerImpl struct {
	MovieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &MovieHandlerImpl{MovieService: movieService}
}

func (h *MovieHandlerImpl) AddMovie(srv echo.Context) error {
	AddMovieRequest := modelsrequest.CreateMovie{}
	err := srv.Bind(&AddMovieRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.MovieService.AddMovie(srv, AddMovieRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Add Movie Error"))
	}

	response := res.CreateMovieToMovieResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Add Movie Data", response))
}

func (h *MovieHandlerImpl) FindAll(srv echo.Context) error {
	result, err := h.MovieService.FindAll(srv)
	if err != nil {
		if strings.Contains(err.Error(), "movies not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Movies Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find All Movies Data Error"))
	}

	response := res.ConvertMovieResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Find All Movie Data", response))
}

func (h *MovieHandlerImpl) FindMovieByID(srv echo.Context) error {
	movieId := srv.Param("id")
	movieIdInt, err := strconv.Atoi(movieId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := h.MovieService.FindMovieByID(srv, movieIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Movie Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find Movie Data Error"))
	}

	response := res.CreateMovieToMovieResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Find Movie Data", response))
}

func (h *MovieHandlerImpl) FindMovieByTitle(srv echo.Context) error {
	movieTitle := srv.Param("title")

	result, err := h.MovieService.FindByTitle(srv, movieTitle)

	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Movie Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Find Movie Data By Title Error"))
	}

	response := res.MovieToMovieResponse(result)

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Find Movie Data", response))
}

func (h *MovieHandlerImpl) UpdateMovie(srv echo.Context) error {
	movieId := srv.Param("id")
	movieIdInt, err := strconv.Atoi(movieId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	MovieUpdateRequest := modelsrequest.UpdateMovie{}
	err = srv.Bind(&MovieUpdateRequest)
	if err != nil {
		return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := h.MovieService.UpdateMovie(srv, MovieUpdateRequest, movieIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return srv.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "admin not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.UpdateMovieToMovieResponse(result)
	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Admin Data", response))
}

func (h *MovieHandlerImpl) DeleteMovie(srv echo.Context) error {
	movieId := srv.Param("id")
	movieIdInt, err := strconv.Atoi(movieId)
	if err != nil {
		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = h.MovieService.DeleteMovie(srv, movieIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			return srv.JSON(http.StatusNotFound, helper.ErrorResponse("movie Not Found"))
		}

		return srv.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Movie Data Error"))
	}

	return srv.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Movie Data", nil))
}