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

type MovieService interface {
	AddMovie(srv echo.Context, request modelsrequest.CreateMovie) (*domain.Movie, error)
	FindAll(srv echo.Context) ([]domain.Movie, error)
	FindMovieByID(srv echo.Context, id int) (*domain.Movie, error)
	FindByTitle(ctx echo.Context, title string) (*domain.Movie, error)
	UpdateMovie(srv echo.Context, request modelsrequest.UpdateMovie, id int) (*domain.Movie, error)
	DeleteMovie(srv echo.Context, id int) error
}

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	Validate        *validator.Validate
}

func NewMovieService(MovieRepository repository.MovieRepository, Validate *validator.Validate) MovieService {
	return &MovieServiceImpl{
		MovieRepository: MovieRepository,
		Validate:        Validate,
	}
}

func (service *MovieServiceImpl) AddMovie(srv echo.Context, request modelsrequest.CreateMovie) (*domain.Movie, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	movie := req.MovieCreateRequestToMovieDomain(&request)

	result, err := service.MovieRepository.AddMovie(movie)
	if err != nil {
		return nil, fmt.Errorf("error adding movie: %s", err.Error())
	}

	return result, nil
}

func (service *MovieServiceImpl) FindAll(srv echo.Context) ([]domain.Movie, error) {
	movie, err := service.MovieRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("movies not found")
	}

	return movie, nil
}

func (service *MovieServiceImpl) FindMovieByID(srv echo.Context, id int) (*domain.Movie, error) {
	movie, err := service.MovieRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (service *MovieServiceImpl) FindByTitle(srv echo.Context, title string) (*domain.Movie, error) {
	movie, _ := service.MovieRepository.FindByTitle(title)
	if movie == nil {
		return nil, fmt.Errorf("movie not found")
	}

	return movie, nil
}

func (service *MovieServiceImpl) UpdateMovie(srv echo.Context, request modelsrequest.UpdateMovie, id int) (*domain.Movie, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helper.ValidationError(srv, err)
	}

	existingMovie, _ := service.MovieRepository.FindByID(id)
	if existingMovie == nil {
		return nil, fmt.Errorf("movie not found")
	}

	movie := req.MovieUpdateRequestToMovieDomain(&request)

	_ , err = service.MovieRepository.Update(movie, id)
	if err != nil {
		return nil, fmt.Errorf("error adding movie: %s", err.Error())
	}

	result, _ := service.MovieRepository.FindByID(id)
	
	return result, nil
}

func (service *MovieServiceImpl) DeleteMovie(srv echo.Context, id int) error {

	existingAdmin, _ := service.MovieRepository.FindByID(id)
	if existingAdmin == nil {
		return fmt.Errorf("movie not found")
	}

	err := service.MovieRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting movie: %s", err)
	}

	return nil
}
