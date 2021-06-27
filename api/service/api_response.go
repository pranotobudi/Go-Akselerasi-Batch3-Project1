package service

import (
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
)

type ResponseUser struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	RoleID    uint   `json:"role_id"`
	AuthToken string `json:"auth_token"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type ResponseGenre struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type ResponseMovie struct {
	ID           uint                  `json:"id"`
	UserID       uint                  `json:"user_id"`
	Title        string                `json:"title"`
	Year         int                   `json:"year"`
	Rating       int                   `json:"rating"`
	MovieReviews []ResponseMovieReview `json:"movie_reviews"`
	Genres       []ResponseGenre       `json:"genres"`
	// CreatedAt    time.Time            `json:"created_at"`
	// UpdatedAt    time.Time            `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt       `json:"deleted_at"`
}

type ResponseMovieJoin struct {
	ID           uint                  `json:"id"`
	UserID       uint                  `json:"user_id"`
	Title        string                `json:"title"`
	Year         int                   `json:"year"`
	Rating       int                   `json:"rating"`
	MovieReviews []ResponseMovieReview `json:"movie_reviews"`
	Genres       []ResponseGenre       `json:"genres"`
	// CreatedAt    time.Time            `json:"created_at"`
	// UpdatedAt    time.Time            `json:"updated_at"`
	// DeletedAt    gorm.DeletedAt       `json:"deleted_at"`
}
type ResponseGenreMovie struct {
	ID        uint   `json:"id"`
	GenreID   uint   `json:"genre_id"`
	MovieID   uint   `json:"movie_id"`
	AuthToken string `json:"auth_token"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type ResponseMovieReview struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	MovieID uint   `json:"movie_id"`
	Review  string `json:"review"`
	Rate    int    `json:"rate"`
	// CreatedAt time.Time      `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func UserResponseFormatter(user entity.User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		AuthToken: auth_token,
		// CreatedAt: user.CreatedAt,
		// UpdatedAt: user.UpdatedAt,
		// DeletedAt: user.DeletedAt,
	}
	return formatter
}

func GenreResponseFormatter(genre entity.Genre) ResponseGenre {
	formatter := ResponseGenre{
		ID:     genre.ID,
		UserID: genre.UserID,
		Name:   genre.Name,
		// CreatedAt: genre.CreatedAt,
		// UpdatedAt: genre.UpdatedAt,
		// DeletedAt: genre.DeletedAt,
	}
	return formatter
}

func MovieResponseFormatter(movie entity.Movie, genres []entity.Genre) ResponseMovie {
	var responseMovieReviews []ResponseMovieReview
	var responseGenres []ResponseGenre
	for _, movieReview := range movie.MovieReviews {
		responseMovieReview := MovieReviewResponseFormatter(movieReview)
		responseMovieReviews = append(responseMovieReviews, responseMovieReview)
	}
	for _, genre := range genres {
		responseGenre := GenreResponseFormatter(genre)
		responseGenres = append(responseGenres, responseGenre)
	}
	formatter := ResponseMovie{
		ID:           movie.ID,
		UserID:       movie.UserID,
		Title:        movie.Title,
		Year:         movie.Year,
		Rating:       movie.Rating,
		MovieReviews: responseMovieReviews,
		Genres:       responseGenres,
		// CreatedAt:    movie.CreatedAt,
		// UpdatedAt:    movie.UpdatedAt,
		// DeletedAt:    movie.DeletedAt,
	}
	return formatter
}

func GenreMovieResponseFormatter(genreMovie entity.GenreMovie) ResponseGenreMovie {
	formatter := ResponseGenreMovie{
		ID:      genreMovie.ID,
		GenreID: genreMovie.GenreID,
		MovieID: genreMovie.MovieID,
		// CreatedAt: genreMovie.CreatedAt,
		// UpdatedAt: genreMovie.UpdatedAt,
		// DeletedAt: genreMovie.DeletedAt,
	}
	return formatter
}

func MovieReviewResponseFormatter(movieReview entity.MovieReview) ResponseMovieReview {
	formatter := ResponseMovieReview{

		ID:      movieReview.ID,
		UserID:  movieReview.UserID,
		MovieID: movieReview.MovieID,
		Review:  movieReview.Review,
		Rate:    movieReview.Rate,
		// CreatedAt: movieReview.CreatedAt,
		// UpdatedAt: movieReview.UpdatedAt,
		// DeletedAt: movieReview.DeletedAt,
	}
	return formatter
}
