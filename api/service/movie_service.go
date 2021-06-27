package service

import (
	"errors"
	"fmt"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/entity"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/repository"
)

type MovieServices interface {
	CheckExistsGenre(req RequestGenre) error
	AddGenre(req RequestGenre) (entity.Genre, error)
	GetGenre(req RequestGenre) (entity.Genre, error)
	GetAllGenres() ([]entity.Genre, error)
	GetAllGenresByMovieID(movieID uint) ([]entity.Genre, error)
	CheckExistsMovie(req RequestMovie) error
	AddMovie(req RequestMovie) (entity.Movie, error)
	GetMovieByID(id uint) (*entity.Movie, error)
	GetAllMovies() ([]entity.Movie, error)
	CheckExistsGenreMovie(req RequestGenreMovie) error
	AddGenreMovie(req RequestGenreMovie) (entity.GenreMovie, error)
	GetGenreMovie(req RequestGenreMovie) (*entity.GenreMovie, error)
	AddMovieReview(req RequestMovieReview) (entity.MovieReview, error)
	GetMovieReview(id uint) ([]entity.MovieReview, error)
}

type movieServices struct {
	repository repository.Repository
}

func NewMovieServices(repository repository.Repository) *movieServices {
	return &movieServices{repository}
}

func (s *movieServices) CheckExistsGenre(req RequestGenre) error {
	name := req.Name
	if genre := s.repository.FindGenre(name); genre != nil {
		return errors.New("genre already inserted")
	}
	return nil
}

func (s *movieServices) AddGenre(req RequestGenre) (entity.Genre, error) {
	genre := entity.Genre{}
	genre.UserID = req.UserID
	genre.Name = req.Name
	newGenre, err := s.repository.AddGenre(genre)
	if err != nil {
		return newGenre, err
	}
	return newGenre, nil
}

func (s *movieServices) GetGenre(req RequestGenre) (entity.Genre, error) {
	var genre = entity.Genre{}
	return genre, nil
}

func (s *movieServices) GetAllGenres() ([]entity.Genre, error) {
	fmt.Println("===========MOVIE-SERVICES: GET-ALL-GENRES==============")
	genres, err := s.repository.GetAllGenres()
	if err != nil {
		return genres, err
	}
	return genres, nil
}
func (s *movieServices) GetAllGenresByMovieID(movieID uint) ([]entity.Genre, error) {
	genres, err := s.repository.GetAllGenresByMovieID(movieID)
	if err != nil {
		return genres, err
	}
	return genres, nil
}
func (s *movieServices) CheckExistsMovie(req RequestMovie) error {
	title := req.Title
	if movie := s.repository.FindMovie(title); movie != nil {
		return errors.New("genre already inserted")
	}
	return nil
}
func (s *movieServices) AddMovie(req RequestMovie) (entity.Movie, error) {
	movie := entity.Movie{}
	movie.Title = req.Title
	movie.Year = req.Year
	movie.Rating = req.Rating

	newMovie, err := s.repository.AddMovie(movie)
	if err != nil {
		return newMovie, err
	}
	return newMovie, nil

}
func (s *movieServices) GetMovieByID(id uint) (*entity.Movie, error) {
	newMovie, err := s.repository.GetMovieByID(id)
	if err != nil {
		return newMovie, err
	}
	return newMovie, nil
}
func (s *movieServices) GetAllMovies() ([]entity.Movie, error) {
	movies, err := s.repository.GetAllMovies()
	// fmt.Printf("\n movieServices GetAllMovies: %+v \n", movies)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (s *movieServices) CheckExistsGenreMovie(req RequestGenreMovie) error {
	genreID := req.GenreID
	movieID := req.MovieID
	if genreMovie := s.repository.FindGenreMovie(genreID, movieID); genreMovie != nil {
		return errors.New("genre for this Movie already inserted")
	}
	return nil
}

func (s *movieServices) AddGenreMovie(req RequestGenreMovie) (entity.GenreMovie, error) {
	genreMovie := entity.GenreMovie{}
	genreMovie.GenreID = req.GenreID
	genreMovie.MovieID = req.MovieID

	newGenreMovie, err := s.repository.AddGenreMovie(genreMovie)
	if err != nil {
		return newGenreMovie, err
	}
	return newGenreMovie, nil

}

func (s *movieServices) GetGenreMovie(req RequestGenreMovie) (*entity.GenreMovie, error) {
	genreMovie := s.repository.FindGenreMovie(req.GenreID, req.MovieID)
	if genreMovie != nil {
		return genreMovie, errors.New("genre for this Movie already inserted")
	}
	return genreMovie, nil

}

func (s *movieServices) AddMovieReview(req RequestMovieReview) (entity.MovieReview, error) {
	movieReview := entity.MovieReview{}
	movieReview.UserID = req.UserID
	movieReview.MovieID = req.MovieID
	movieReview.Rate = req.Rate
	movieReview.Review = req.Review

	newMovieReview, err := s.repository.AddMovieReview(movieReview)
	if err != nil {
		return newMovieReview, err
	}
	return newMovieReview, nil

}

func (s *movieServices) GetMovieReview(movieID uint) ([]entity.MovieReview, error) {
	movieReviews, err := s.repository.GetMovieReview(movieID)
	if err != nil {
		return movieReviews, err
	}
	return movieReviews, nil
}
