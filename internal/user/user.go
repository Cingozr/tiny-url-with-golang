package user

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	UrlLimit uint
}

type UserService struct {
	DB *gorm.DB
}

type IUserService interface {
	Create(user User) (User, error)
	GetUser(user User) error
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (u *UserService) Create(user User) (User, error) {
	if result := u.DB.Save(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (u *UserService) GetUser(user User) error {
	var userModel User
	if result := u.DB.Find(&userModel).Where("email = ? and password = ?", user.Email, user.Password); result.Error != nil {
		return result.Error
	}

	if userModel.Email == "" {
		return errors.New("couldn't find user in db")
	}
	return nil
}
