package request

import (
	"Mini-Project/models/domain"
	modelsrequest "Mini-Project/models/models-request"
)

func MapCreateStudioRequestToStudio(req *modelsrequest.CreateStudioRequest) *domain.Studio {
	return &domain.Studio{
		Name: req.Name,
	}
}
