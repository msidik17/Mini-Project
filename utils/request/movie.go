package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/models/schema"
)

func MovieDomaintoMovieSchema(movie domain.Movie) *schema.Movie {
	return &schema.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Studio:      movie.Studio,
		Price:       movie.Price,
	}
}

func MovieCreateRequestToMovieDomain(request *modelsrequest.CreateMovie) *domain.Movie {
	return &domain.Movie{
		Title:       request.Title,
		Description: request.Description,
		Studio:      request.Studio,
		Price:       request.Price,
	}
}

func MovieUpdateRequestToMovieDomain(request *modelsrequest.UpdateMovie) *domain.Movie {
	return &domain.Movie{
		Title:       request.Title,
		Description: request.Description,
		Studio:      request.Studio,
		Price:       request.Price,
	}
}
