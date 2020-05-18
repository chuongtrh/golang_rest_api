package auth

import (
	"demo_api/src/config"
	"demo_api/src/modules/dto"
	"demo_api/src/modules/user"
	"demo_api/src/utils"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	userService user.Service
}

func NewAuthController(userService user.Service) Controller {
	return Controller{
		userService: userService,
	}
}

func (controller Controller) Login(c echo.Context) error {
	loginDTO := new(dto.LoginDTO)
	if err := c.Bind(loginDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if errs := c.Validate(loginDTO); errs != nil {
		log.Error().Msg(errs.Error())
		return c.JSON(http.StatusBadRequest, errs.Error())
	}

	user, err := controller.userService.Login(loginDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	token, refreshToken, err := utils.CreateAuthTokenPair(user.Email, user.ID, user.Role, config.Cfg.JwtKey, config.Cfg.JwtExp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userInfo":     user.ToObject(),
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (controller Controller) Refresh(c echo.Context) error {
	log.Info().Msg("Refresh")
	return c.JSON(http.StatusOK, "Refresh")
}
