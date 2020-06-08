package user

import (
	"demo_api/src/util"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Controller interface
type Controller interface {
	GetMyProfile(c echo.Context) error
	GetAll(c echo.Context) error
	GetUser(c echo.Context) error
}

// NewUserController func
func NewUserController(userService Service) (Controller, error) {
	return &controller{
		userService: userService,
	}, nil
}

type controller struct {
	userService Service
}

func (controller *controller) GetMyProfile(c echo.Context) error {
	claims := c.Get("user").(*util.Claims)
	//log.Info().Msgf("claims:%+v", claims)

	id := uint64(claims.ID)
	user, err := controller.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}
	return c.JSON(http.StatusOK, user.ToObject())
}

func (controller *controller) GetAll(c echo.Context) error {
	users, err := controller.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *controller) GetUser(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}

	id := uint64(idP)
	user, err := controller.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errors.New("Not found"))
	}
	return c.JSON(http.StatusOK, user.ToObject())
}
