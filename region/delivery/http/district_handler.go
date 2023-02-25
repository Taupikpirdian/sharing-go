package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/region/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func isRequestValid(m *domain.DistrictData) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register will store the user by given request body
func (a *RegionHandler) DistrictGet(c echo.Context) (err error) {
	var district domain.DistrictData
	err = c.Bind(&district)
	if err != nil {
		responseError, _ := http_response.MapResponseDistrict(1, err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid(&district); !ok {
		responseError2, _ := http_response.MapResponseDistrict(1, err.Error(), nil)
		return c.JSON(http.StatusBadRequest, responseError2)
	}

	ctx := c.Request().Context()
	data, errUc := a.AUsecase.GetDistrict(ctx)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseDistrict(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseDistrict(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
