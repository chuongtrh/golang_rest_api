package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// type Controller interface {
// 	GetAll(c echo.Context) error
// 	GetUser(c echo.Context) error
// }

type Controller struct {
	userService Service
}

func NewUserController(userService Service) Controller {
	return Controller{
		userService: userService,
	}
}

func (controller Controller) GetMyProfile(c echo.Context) error {
	return c.JSON(http.StatusOK, "get my profile")
}

func (controller Controller) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, "get users")
}

func (controller Controller) GetUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}

	id := int64(idP)
	user, err := controller.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}
	return c.JSON(http.StatusOK, user.ToObject())
}
