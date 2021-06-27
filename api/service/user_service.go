package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServices interface {
	CreateUser(req RequestUser) (entity.User, error)
	CheckExistsEmail(req RequestUser) error
	AuthUser(req RequestUserLogin) (entity.User, error)
	GetRole(userID uint) (string, error)
	GetAllUsers() ([]entity.User, error)
	GetUser(email string) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	GetRegistration(email string) (*entity.Registration, error)
	UpdateUser(req RequestUser) (entity.User, error)
	CheckUserExists(req RequestUser) (bool, error)
	AddRegistrationSendEmail(req RequestUser, regToken string) (entity.Registration, error)
}

type userServices struct {
	repository repository.Repository
}

func NewUserServices(repository repository.Repository) *userServices {
	return &userServices{repository}
}

func (s *userServices) CreateUser(req RequestUser) (entity.User, error) {
	user := entity.User{}
	user.Name = req.Name
	user.Email = req.Email
	user.RoleID = req.RoleID

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

func (s *userServices) CheckUserExists(req RequestUser) (bool, error) {
	email := req.Email
	if user := s.repository.FindEmail(email); user != nil {
		return true, errors.New("user already registered")
	}
	return false, errors.New("user not registered")
}

func (s *userServices) GetRole(userID uint) (string, error) {
	role, err := s.repository.FindRole(userID)
	fmt.Printf("\n Service: userID: %v ROLE: %+v \n", userID, role)
	if err != nil {
		return role, errors.New("role not found")
	}
	return role, nil
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

func (s *userServices) GetAllUsers() ([]entity.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userServices) GetUser(email string) (*entity.User, error) {
	user, err := s.repository.GetUser(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userServices) GetUserByID(id uint) (*entity.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userServices) GetRegistration(email string) (*entity.Registration, error) {
	registration, err := s.repository.GetRegistration(email)
	if err != nil {
		return registration, err
	}
	return registration, nil
}

func (s *userServices) UpdateUser(req RequestUser) (entity.User, error) {
	user := entity.User{}
	user.ID = req.ID
	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password
	user.RoleID = req.RoleID

	newUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *userServices) AddRegistrationSendEmail(req RequestUser, regToken string) (entity.Registration, error) {
	registration := entity.Registration{}
	registration.ID = req.ID
	registration.Email = req.Email
	registration.Name = req.Name
	registration.Password = req.Password
	registration.RoleID = req.RoleID
	registration.RegistrationToken = regToken
	registration.TimeCreated = time.Now()

	registration, err := s.repository.AddRegistration(registration)
	return registration, err
}
