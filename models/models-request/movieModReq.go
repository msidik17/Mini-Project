package modelsrequest

type CreateMovie struct {
	Title       string  `json:"title" form:"title" validate:"required,min=1,max=255"`
	Description string  `json:"description" form:"description" validate:"required,min=1,max=255"`
	Studio      string  `json:"studio" form:"studio" validate:"required,min=1,max=255"`
	Price       float64 `json:"price" form:"price" validate:"required,min=1"`
}

type UpdateMovie struct {
	Title       string  `json:"title" form:"title" validate:"required,min=1,max=255"`
	Description string  `json:"description" form:"description" validate:"required,min=1,max=255"`
	Studio      string  `json:"studio" form:"studio" validate:"required,min=1,max=255"`
	Price       float64 `json:"price" form:"price" validate:"required,min=1"`
}
