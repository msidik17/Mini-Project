package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
	"Mini-Project/models/schema"
)

func UserDomainToUserLoginResponse(user *domain.User) modelsresponse.UserLoginResponse {
	return modelsresponse.UserLoginResponse{
		Name:  user.Name,
		Email: user.Email,
	}
}

func UserSchemaToUserDomain(user *schema.User) *domain.User {
	return &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func UserDomaintoUserResponse(user *domain.User) modelsresponse.UserReponse {
	return modelsresponse.UserReponse{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ConvertUserResponse(users []domain.User) []modelsresponse.UserReponse {
	var results []modelsresponse.UserReponse
	for _, user := range users {
		userResponse := modelsresponse.UserReponse{
			Id:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		}
		results = append(results, userResponse)
	}
	return results
}
