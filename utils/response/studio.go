package response

import (
	"Mini-Project/models/domain"
	modelsresponse "Mini-Project/models/models-response"
)

func MapStudioToStudioResponse(studio *domain.Studio) *modelsresponse.StudioResponse {
	return &modelsresponse.StudioResponse{
		ID:   studio.ID,
		Name: studio.Name,
	}
}
