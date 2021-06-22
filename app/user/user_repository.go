package user

import "gorm.io/gorm"

type UsersStorage []interface{}

type Repository interface {
	InsertUser(user User) (User, error)
	FindEmail(email string) *User
	FindUserByEmail(email string) (User, error)
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
func (r *repository) InsertUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	// users = append(users, user)
	return user, nil
}

func (r *repository) FindEmail(email string) *User {
	var user User
	err := r.db.First(&user, "email=?", email).Error
	if err == nil {
		return &user
	}
	return nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User
	// err := r.db.Where("email = ?", email).Find(&user).Error
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
