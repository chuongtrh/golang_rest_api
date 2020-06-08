package user

import (
	"demo_api/src/dto"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service interface
type Service interface {
	GetAll() ([]User, error)
	GetUser(id uint64) (User, error)
	Login(dto *dto.LoginDTO) (User, error)
}

// NewUserService func
func NewUserService(repository Repository) (Service, error) {
	return &service{
		repository: repository,
	}, nil
}

type service struct {
	repository Repository
}

func (service *service) GetAll() ([]User, error) {
	return service.repository.GetAll()
}
func (service *service) GetUser(id uint64) (User, error) {
	return service.repository.GetUser(id)
}

func (service *service) Login(dto *dto.LoginDTO) (User, error) {
	user, err := service.repository.GetUserByEmail(dto.Email)
	if err != nil {
		return User{}, errors.New("wrong email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		return User{}, errors.New("wrong password")
	}
	return user, nil
}
