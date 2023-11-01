package modelsresponse

type MovieResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Studio      string  `json:"studio"`
	Price       float64 `json:"price"`
}

type CreateMovieResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Studio      string  `json:"studio"`
	Price       float64 `json:"price"`
}

type UpdateMovieResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Studio      string  `json:"studio"`
	Price       float64 `json:"price"`
}
