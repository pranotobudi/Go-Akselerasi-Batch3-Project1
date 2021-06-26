package repository

import (
	"errors"
	"fmt"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"gorm.io/gorm"
)

type UsersStorage []interface{}

type Repository interface {
	AddUser(user entity.User) (entity.User, error)
	FindEmail(email string) *entity.User
	GetUserByEmail(email string) (entity.User, error)
	FindGenre(name string) *entity.Genre
	AddGenre(genre entity.Genre) (entity.Genre, error)
	GetAllGenres() ([]entity.Genre, error)
	FindMovie(name string) *entity.Movie
	AddMovie(movie entity.Movie) (entity.Movie, error)
	GetAllMovies() ([]entity.Movie, error)
	FindGenreMovie(genreID uint, movieID uint) *entity.GenreMovie
	AddGenreMovie(genreMovie entity.GenreMovie) (entity.GenreMovie, error)
	GetMovieReview(id uint) (*entity.MovieReview, error)
	AddMovieReview(movieReview entity.MovieReview) (entity.MovieReview, error)
	FindRole(userID uint) (string, error)
	GetAllUsers() ([]entity.User, error)
	GetUser(email string) (*entity.User, error)
	GetRegistration(email string) (*entity.Registration, error)
	UpdateUser(newUser entity.User) (entity.User, error)
	AddRegistration(registration entity.Registration) (entity.Registration, error)
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
func (r *repository) AddUser(user entity.User) (entity.User, error) {
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

func (r *repository) FindGenre(name string) *entity.Genre {
	var genre entity.Genre
	err := r.db.First(&genre, "name=?", name).Error
	if err == nil {
		return &genre
	}
	return nil
}
func (r *repository) AddGenre(genre entity.Genre) (entity.Genre, error) {
	err := r.db.Create(&genre).Error
	if err != nil {
		return genre, err
	}
	// users = append(users, user)
	return genre, nil
}

func (r *repository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	// err := r.db.Where("email = ?", email).Find(&user).Error
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) GetAllGenres() ([]entity.Genre, error) {
	fmt.Println("===========REPOSITORY: GET-ALL-GENRES==============")
	var genres []entity.Genre

	result := r.db.Find(&genres)
	fmt.Println("===========REPOSITORY: GET-ALL-GENRES - After db.Find==============")
	fmt.Println("=============len: ", len(genres))
	if result.Error != nil {
		fmt.Println("===========REPOSITORY: GET-ALL-GENRES - inside error != nil ==============")
		return genres, result.Error
	} else if result.RowsAffected < 1 {
		// return genres, fmt.Errorf("table is empty")
		return genres, errors.New("table is empty")
	}
	return genres, nil

}

func (r *repository) FindMovie(title string) *entity.Movie {
	var movie entity.Movie
	err := r.db.First(&movie, "title=?", title).Error
	if err == nil {
		return &movie
	}
	return nil
}
func (r *repository) AddMovie(movie entity.Movie) (entity.Movie, error) {
	err := r.db.Create(&movie).Error
	if err != nil {
		return movie, err
	}
	// users = append(users, user)
	return movie, nil
}

func (r *repository) GetAllMovies() ([]entity.Movie, error) {
	var movies []entity.Movie
	result := r.db.Find(&movies)
	if result.Error != nil {
		return movies, result.Error
	} else if result.RowsAffected < 1 {
		return movies, fmt.Errorf("table is empty")
	}
	return movies, nil
}
func (r *repository) FindGenreMovie(genreID uint, movieID uint) *entity.GenreMovie {
	var genreMovie entity.GenreMovie
	err := r.db.First(&genreMovie, "genre_id=? AND movie_id=?", genreID, movieID).Error
	if err == nil {
		return &genreMovie
	}
	return nil

}

func (r *repository) AddGenreMovie(genreMovie entity.GenreMovie) (entity.GenreMovie, error) {
	err := r.db.Create(&genreMovie).Error
	if err != nil {
		return genreMovie, err
	}
	// users = append(users, user)
	return genreMovie, nil
}

func (r *repository) AddMovieReview(movieReview entity.MovieReview) (entity.MovieReview, error) {
	err := r.db.Create(&movieReview).Error
	if err != nil {
		return movieReview, err
	}
	// users = append(users, user)
	return movieReview, nil
}

func (r *repository) GetMovieReview(id uint) (*entity.MovieReview, error) {
	var movieReview entity.MovieReview
	err := r.db.First(&movieReview, "id=?", id).Error
	if err != nil {
		return &movieReview, err
	}
	return &movieReview, nil
}

func (r *repository) FindRole(userID uint) (string, error) {
	var role string
	statement := "SELECT roles.role FROM users JOIN roles ON users.role_id=roles.id WHERE users.id = ?"
	//take tomorrow events which happen for the next 23-24 hours.
	result := r.db.Raw(statement, userID).Find(&role)
	fmt.Printf("\n Repository: userID: %v ROLE: %+v \n", userID, role)
	return role, result.Error
}

func (r *repository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	} else if result.RowsAffected < 1 {
		return users, fmt.Errorf("table is empty")
	}
	return users, nil
}

func (r *repository) GetUser(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}
func (r *repository) GetRegistration(email string) (*entity.Registration, error) {
	var registration entity.Registration
	err := r.db.First(&registration, "email=?", email).Error
	if err != nil {
		return &registration, err
	}
	return &registration, nil
}

func (r *repository) UpdateUser(newUser entity.User) (entity.User, error) {
	err := r.db.Save(&newUser).Error
	if err != nil {
		return newUser, err
	}
	// users = append(users, user)
	return newUser, nil
}

func (r *repository) AddRegistration(registration entity.Registration) (entity.Registration, error) {
	err := r.db.Create(&registration).Error
	if err != nil {
		return registration, err
	}
	return registration, nil

}
