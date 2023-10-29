package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

func MovieToMovieResponse(movie *domain.Movie) modelsresponse.MovieResponse {
	return modelsresponse.MovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
	}
}

func CreateMovieToMovieResponse(movie *domain.Movie) modelsresponse.CreateMovieResponse {
	return modelsresponse.CreateMovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
	}
}

func UpdateMovieToMovieResponse(movie *domain.Movie) modelsresponse.UpdateMovieResponse {
	return modelsresponse.UpdateMovieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
	}
}

func MovieSchematoMovieDomain(movie *schema.Movie) *domain.Movie {
	return &domain.Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
	}
}


func ConvertMovieResponse(movies []domain.Movie) []modelsresponse.MovieResponse {
	var results []modelsresponse.MovieResponse
	for _, movie := range movies {
		movieResponse := modelsresponse.MovieResponse{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
		}
		results = append(results, movieResponse)
	}
	return results
}
