package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

func AdminDomainToAdminLoginResponse(user *domain.Admin) modelsresponse.AdminLoginResponse {
	return modelsresponse.AdminLoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func AdminSchemaToAdminDomain(user *schema.Admin) *domain.Admin {
	return &domain.Admin{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func AdminDomaintoAdminResponse(user *domain.Admin) modelsresponse.AdminReponse {
	return modelsresponse.AdminReponse{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
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
