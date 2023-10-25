package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/models/schema"
)

func UserCreateRequestToUserDomain(request modelsrequest.UserCreate) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserLoginRequestToUserDomain(request modelsrequest.UserLogin) *domain.User {
	return &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserUpdateRequestToUserDomain(request modelsrequest.UserUpdate) *domain.User {
	return &domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func UserDomainToUserSchema(request domain.User) *schema.User {
	return &schema.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
