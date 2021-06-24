package service

import (
	"time"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"gorm.io/gorm"
)

type ResponseUser struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	AuthToken string         `json:"auth_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type ResponseGenre struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	Name      string         `json:"name"`
	AuthToken string         `json:"auth_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type ResponseMovie struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	Title     string         `json:"title"`
	Year      int            `json:"year"`
	Rating    int            `json:"rating"`
	AuthToken string         `json:"auth_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ResponseGenreMovie struct {
	ID        uint           `json:"id"`
	GenreID   uint           `json:"genre_id"`
	MovieID   uint           `json:"movie_id"`
	AuthToken string         `json:"auth_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ResponseMovieReview struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	MovieID   uint           `json:"movie_id"`
	Review    string         `json:"review"`
	Rate      int            `json:"rate"`
	AuthToken string         `json:"auth_token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func UserResponseFormatter(user entity.User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		Name:      user.Name,
		Email:     user.Email,
		AuthToken: auth_token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
	return formatter
}

func GenreResponseFormatter(genre entity.Genre, auth_token string) ResponseGenre {
	formatter := ResponseGenre{
		ID:        genre.ID,
		UserID:    genre.UserID,
		Name:      genre.Name,
		AuthToken: auth_token,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
		DeletedAt: genre.DeletedAt,
	}
	return formatter
}

func MovieResponseFormatter(movie entity.Movie, auth_token string) ResponseMovie {
	formatter := ResponseMovie{
		ID:        movie.ID,
		UserID:    movie.UserID,
		Title:     movie.Title,
		Year:      movie.Year,
		Rating:    movie.Rating,
		AuthToken: auth_token,
		CreatedAt: movie.CreatedAt,
		UpdatedAt: movie.UpdatedAt,
		DeletedAt: movie.DeletedAt,
	}
	return formatter
}

func GenreMovieResponseFormatter(genreMovie entity.GenreMovie, auth_token string) ResponseGenreMovie {
	formatter := ResponseGenreMovie{
		ID:        genreMovie.ID,
		GenreID:   genreMovie.GenreID,
		MovieID:   genreMovie.MovieID,
		AuthToken: auth_token,
		CreatedAt: genreMovie.CreatedAt,
		UpdatedAt: genreMovie.UpdatedAt,
		DeletedAt: genreMovie.DeletedAt,
	}
	return formatter
}

func MovieReviewResponseFormatter(movieReview entity.MovieReview, auth_token string) ResponseMovieReview {
	formatter := ResponseMovieReview{

		ID:        movieReview.ID,
		UserID:    movieReview.UserID,
		MovieID:   movieReview.MovieID,
		Review:    movieReview.Review,
		Rate:      movieReview.Rate,
		AuthToken: auth_token,
		CreatedAt: movieReview.CreatedAt,
		UpdatedAt: movieReview.UpdatedAt,
		DeletedAt: movieReview.DeletedAt,
	}
	return formatter
}
