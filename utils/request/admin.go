package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
	"Mini-Project/models/schema"
)

func AdminCreateRequestToAdminDomain(request modelsrequest.AdminCreate) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request modelsrequest.AdminLogin) *domain.Admin {
	return &domain.Admin{
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request modelsrequest.AdminUpdate) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminDomainToAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
