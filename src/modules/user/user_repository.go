package user

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {

	return Repository{
		db: db,
	}
}

func (repo Repository) GetUser(id int64) (User, error) {
	var err error
	var u User
	err = repo.db.Raw("SELECT * FROM User WHERE id = ?", id).Scan(&u).Error
	if err != nil {
		return u, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return u, errors.New("user not found")
	}
	return u, err
}

func (repo Repository) GetUserByEmail(email string) (User, error) {
	var err error
	var u User
	err = repo.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&u).Error
	if err != nil {
		return u, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return u, errors.New("user not found")
	}
	return u, err
}

func (repo Repository) GetAll() ([]User, error) {
	var err error
	var users []User
	err = repo.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, err
}
