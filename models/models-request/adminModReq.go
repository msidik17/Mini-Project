package modelsrequest

type AdminCreate struct {
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

type AdminLogin struct {
	Email string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}

type AdminUpdate struct{
	Name string `json:"name" form:"name" validate:"required,min=1,max=255"`
	Email string `json:"email" form:"email" validate:"required,min=1,max=255"`
	Password string `json:"password" form:"password" validate:"required,min=1,max=255"`
}