package http

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/user/delivery/http_response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Register will store the user by given request body
func (a *UserHandler) Register(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		responseError, _ := http_response.MapResponse(1, err.Error())
		return c.JSON(http.StatusUnprocessableEntity, responseError)
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		responseErrorRequest, _ := http_response.MapResponse(1, domain.ErrBadParamInput.Error())
		return c.JSON(http.StatusBadRequest, responseErrorRequest)
	}

	ctx := c.Request().Context()
	err = a.AUsecase.Register(ctx, &user)
	if err != nil {
		responseErrorUsecase, _ := http_response.MapResponse(1, domain.ErrBadBody.Error())
		return c.JSON(getStatusCode(err), responseErrorUsecase)
	}

	responseSuccess, _ := http_response.MapResponse(0, "success")
	return c.JSON(http.StatusCreated, responseSuccess)
}
