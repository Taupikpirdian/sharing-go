package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler  represent the httphandler for user
type RegionHandler struct {
	AUsecase domain.RegionUsecase
}

// NewUserHandler will initialize the users/ resources endpoint
func NewRegionHandler(e *echo.Echo, us domain.RegionUsecase) {
	handler := &RegionHandler{
		AUsecase: us,
	}
	e.GET("/api/v1/district", handler.DistrictGet)
	e.GET("/api/v1/village", handler.VillageGet)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
