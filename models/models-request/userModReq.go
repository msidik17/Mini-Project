package modelsrequest

type UserCreate struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

type UserUpdate struct {
	Name     string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email    string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}
