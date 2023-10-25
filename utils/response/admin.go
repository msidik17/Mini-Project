package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

func AdminDomainToAdminLoginResponse(admin *domain.Admin) modelsresponse.AdminLoginResponse {
	return modelsresponse.AdminLoginResponse{
		Name: admin.Name,
		Email: admin.Email,
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

func ConvertAdminResponse(users []domain.Admin) []modelsresponse.AdminReponse {
	var results []modelsresponse.AdminReponse
	for _, user := range users {
		adminResponse := modelsresponse.AdminReponse{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		results = append(results, adminResponse)
	}
	return results
}
