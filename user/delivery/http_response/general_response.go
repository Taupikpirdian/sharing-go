package http_response

import "github.com/bxcodec/go-clean-arch/domain"

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
	Data   []*domain.Job
}

type ResponseItemJson struct {
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

func MapResponse(code int, message string) (*CustomReponseSingle, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	return httpResponse, nil
}

func MapResponseJob(code int, message string, data []*domain.Job) (*CustomReponseSingle, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	return httpResponse, nil
}
