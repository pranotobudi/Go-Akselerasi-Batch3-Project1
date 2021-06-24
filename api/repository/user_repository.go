package repository

import (
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"gorm.io/gorm"
)

type UsersStorage []interface{}

type Repository interface {
	InsertUser(user entity.User) (entity.User, error)
	FindEmail(email string) *entity.User
	FindUserByEmail(email string) (entity.User, error)
}

// var users UsersStorage
// var users *gorm.DB

type repository struct {
	// users *UsersStorage
	db *gorm.DB
}

// func NewRepository(users *UsersStorage) *repository {
// 	return &repository{users}
// }
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) InsertUser(user User) User {
// 	users = append(users, user)
// 	return user
// }
func (r *repository) InsertUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	// users = append(users, user)
	return user, nil
}

func (r *repository) FindEmail(email string) *entity.User {
	var user entity.User
	err := r.db.First(&user, "email=?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (r *repository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User
	// err := r.db.Where("email = ?", email).Find(&user).Error
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
