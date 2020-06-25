package user

import (
	"demo_api/src/dto"
	"demo_api/src/util"
	"demo_api/src/util/logger"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Controller interface
type Controller interface {
	GetMyProfile(c echo.Context) error
	GetAll(c echo.Context) error
	GetUser(c echo.Context) error
	Register(c echo.Context) error
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
	//logger.Infof("claims:%+v", claims)

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

func (controller *controller) Register(c echo.Context) error {
	registerDTO := new(dto.RegisterDTO)
	if err := c.Bind(registerDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errs := c.Validate(registerDTO); errs != nil {
		logger.Error(errs.Error())
		return c.JSON(http.StatusBadRequest, errs.Error())
	}

	isEmailExisted, err := controller.userService.CheckEmailExist(registerDTO.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if isEmailExisted {
		return c.JSON(http.StatusBadRequest, "Email is existed")
	}

	password := util.HashPassword(registerDTO.Password)
	println("password:", password)

	user, err := controller.userService.Create(registerDTO.Email, password, RoleUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user.ToObject())
}
