package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

func AdminDomainToAdminLoginResponse(admin *domain.Admin) modelsresponse.AdminLoginResponse {
	return modelsresponse.AdminLoginResponse{
		Email: admin.Email,
		Password: admin.Password,
	}
}

func AdminSchemaToAdminDomain(admin *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
}

func AdminDomaintoAdminResponse(admin *domain.Admin) modelsresponse.AdminReponse {
	return modelsresponse.AdminReponse{
		Id:       admin.ID,
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
}

func ConvertAdminResponse(admins []domain.Admin) []modelsresponse.AdminReponse {
	var results []modelsresponse.AdminReponse
	for _, admin := range admins {
		adminResponse := modelsresponse.AdminReponse{
			Id:       admin.ID,
			Name:     admin.Name,
			Email:    admin.Email,
			Password: admin.Password,
		}
		results = append(results, adminResponse)
	}
	return results
}
