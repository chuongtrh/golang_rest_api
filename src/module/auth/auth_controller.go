package auth

import (
	"demo_api/src/config"
	"demo_api/src/dto"
	"demo_api/src/module/user"
	"demo_api/src/util"
	"demo_api/src/util/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller interface
type Controller interface {
	Login(c echo.Context) error
	Refresh(c echo.Context) error
}

type controller struct {
	userService user.Service
}

// NewAuthController func
func NewAuthController(userService user.Service) (Controller, error) {
	return &controller{
		userService: userService,
	}, nil
}

func (controller *controller) Login(c echo.Context) error {
	loginDTO := new(dto.LoginDTO)
	if err := c.Bind(loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if errs := c.Validate(loginDTO); errs != nil {
		logger.Error(errs.Error())
		return c.JSON(http.StatusBadRequest, errs.Error())
	}

	user, err := controller.userService.Login(loginDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, refreshToken, err := util.CreateAuthTokenPair(user.Email, user.ID, user.Role, config.Cfg.JwtKey, config.Cfg.JwtExp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userInfo":     user.ToObject(),
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (controller *controller) Refresh(c echo.Context) error {
	logger.Info("Refresh")
	return c.JSON(http.StatusOK, "Refresh")
}
