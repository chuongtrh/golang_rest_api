package user

import (
	"demo_api/src/modules/dto"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository Repository
}

func NewUserService(userRepository Repository) Service {
	return Service{
		repository: userRepository,
	}
}

func (service Service) GetAll() ([]User, error) {
	return service.repository.GetAll()
}
func (service Service) GetUser(id int64) (User, error) {
	return service.repository.GetUser(id)
}

func (service Service) Login(dto *dto.LoginDTO) (User, error) {
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
