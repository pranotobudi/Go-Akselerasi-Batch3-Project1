package service

import (
	"errors"
	"fmt"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	CreateUser(req RequestUser) (entity.User, error)
	CheckExistsEmail(req RequestUser) error
	AuthUser(req RequestUserLogin) (entity.User, error)
}

type userServices struct {
	repository repository.Repository
}

func NewServices(repository repository.Repository) *userServices {
	return &userServices{repository}
}

func (s *userServices) CreateUser(req RequestUser) (entity.User, error) {
	user := entity.User{}
	user.Name = req.Name
	user.Email = req.Email
	// user.Password = req.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)

	newUser, err := s.repository.AddUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *userServices) CheckExistsEmail(req RequestUser) error {
	email := req.Email
	if user := s.repository.FindEmail(email); user != nil {
		return errors.New("email already registered")
	}
	return nil
}

func (s *userServices) AuthUser(req RequestUserLogin) (entity.User, error) {
	email := req.Email
	password := req.Password
	fmt.Println("AUTHUSER CALLED")

	user, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return user, errors.New("email not registered")
	}
	test, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("COMPARES: %s %s \n", user.Password, string(test))
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}
