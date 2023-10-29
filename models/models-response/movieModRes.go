package modelsresponse

type MovieResponse struct {
    ID          uint    `json:"id" form:"id"`
    Title       string  `json:"title" form:"title"`
    Description string `json:"description" form:"description"`
}

type CreateMovieResponse struct {
    ID          uint    `json:"id" form:"id"`
    Title       string  `json:"title" form:"title"`
    Description string `json:"description" form:"description"`
}

type UpdateMovieResponse struct {
    ID          uint    `json:"id" form:"id"`
    Title       string  `json:"title" form:"title"`
    Description string `json:"description" form:"description"`
}