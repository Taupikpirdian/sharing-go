package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type CustomReponseVillage struct {
	Status *Status
	Data   []*domain.VillageData
}

func MapResponseVillage(code int, message string, data []*domain.VillageData) (*CustomReponseVillage, error) {
	httpResponse := &CustomReponseVillage{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
