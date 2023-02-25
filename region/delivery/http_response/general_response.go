package http_response

import (
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type CustomReponseSingle struct {
	Status *Status
	Data   []*domain.DistrictData
}

type ResponseItemDistrictJson struct {
	Id        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func MapResponseDistrict(code int, message string, dataDistrict []*domain.DistrictData) (*CustomReponseSingle, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: dataDistrict,
	}

	return httpResponse, nil
}
