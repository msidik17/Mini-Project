package repository

import (
	"Mini-Project/models/domain"
	"Mini-Project/models/schema"
	"Mini-Project/utils/request"
	"Mini-Project/utils/response"

	"gorm.io/gorm"
)

type MovieRepository interface {
	AddMovie(movie *domain.Movie) (*domain.Movie, error)
	Update(movie *domain.Movie, Id int) (*domain.Movie, error)
	Delete(id int) error
	FindAll() ([]domain.Movie, error)
	FindByID(id int) (*domain.Movie, error)
	FindByTitle(title string) ([]domain.Movie, error)
}

type MovieRepositoryImpl struct {
	DB *gorm.DB
}

func NewMovieRepository(DB *gorm.DB) MovieRepository {
	return &MovieRepositoryImpl{DB: DB}
}

func (repository *MovieRepositoryImpl) AddMovie(movie *domain.Movie) (*domain.Movie, error) {
	movieDb := request.MovieDomaintoMovieSchema(*movie)
	result := repository.DB.Create(&movieDb)
	if result.Error != nil {
		return nil, result.Error
	}

	results := response.MovieSchematoMovieDomain(movieDb)

	return results, nil
}

func (repository *MovieRepositoryImpl) Update(movie *domain.Movie, id int) (*domain.Movie, error) {
	result := repository.DB.Table("Movies").Where("id = ?", id).Updates(domain.Movie{Title: movie.Title, Description: movie.Description, Studio: movie.Studio, Price: movie.Price})
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}

func (repository *MovieRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Movie{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *MovieRepositoryImpl) FindAll() ([]domain.Movie, error) {
	var movie []domain.Movie

	if err := repository.DB.Where("deleted_at IS NULL").Find(&movie).Error; err != nil {
		return nil, err
	}

	return movie, nil
}

func (repository *MovieRepositoryImpl) FindByID(id int) (*domain.Movie, error) {
	var movie domain.Movie
	if err := repository.DB.Unscoped().Where("id = ? AND deleted_at IS NULL", id).First(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (repository *MovieRepositoryImpl) FindByTitle(title string) ([]domain.Movie, error) {
	var movie []domain.Movie
	result := repository.DB.Unscoped().Where("title =? AND deleted_at IS NULL", title).Find(&movie)
	if result.Error != nil {
		return nil, result.Error
	}
	return movie, nil
}
