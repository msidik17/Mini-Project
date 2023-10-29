package modelsrequest

type CreateMovie struct {
    Title       string `json:"title" form:"title" validate:"required,min=1,max=255"`
    Description string `json:"description" form:"description" validate:"required,min=1,max=255"`
}

type UpdateMovie struct {
    Title       string `json:"title" form:"title" validate:"required,min=1,max=255"`
    Description string `json:"description" form:"description" validate:"required,min=1,max=255"`
}