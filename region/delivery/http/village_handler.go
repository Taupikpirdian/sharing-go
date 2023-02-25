package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/region/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func isRequestValidVillage(m *domain.VillageData) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register will store the user by given request body
func (a *RegionHandler) VillageGet(c echo.Context) (err error) {
	subDistrictId := c.QueryParam("subDistrictId")

	ctx := c.Request().Context()
	data, errUc := a.AUsecase.GetVillage(ctx, subDistrictId)
	if errUc != nil {
		responseError3, _ := http_response.MapResponseVillage(1, domain.ErrBadBody.Error(), nil)
		return c.JSON(getStatusCode(err), responseError3)
	}

	responseSuccess, _ := http_response.MapResponseVillage(0, "success", data)
	return c.JSON(http.StatusOK, responseSuccess)
}
