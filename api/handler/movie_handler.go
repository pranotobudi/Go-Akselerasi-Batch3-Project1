package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
)

type movieHandler struct {
	movieService service.MovieServices
	authService  auth.AuthService
}

func NewMovieHandler(movieService service.MovieServices, authService auth.AuthService) *movieHandler {
	return &movieHandler{movieService, authService}
}

func (h *movieHandler) AddGenre(c echo.Context) error {
	genre := new(service.RequestGenre)
	if err := c.Bind(genre); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	err := h.movieService.CheckExistsGenre(*genre)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newGenre, _ := h.movieService.AddGenre(*genre)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.GenreResponseFormatter(newGenre, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "new genre successfully added", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) GetAllGenres(c echo.Context) error {
	genres, err := h.movieService.GetAllGenres()
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	// auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	var finalUserData []service.ResponseGenre
	for _, genre := range genres {
		auth_token, err := h.authService.GetAccessToken(genre.ID)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

			return c.JSON(http.StatusInternalServerError, response)
		}

		userData := service.GenreResponseFormatter(genre, auth_token)
		finalUserData = append(finalUserData, userData)
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "get all genres succeeded", finalUserData)

	return c.JSON(http.StatusOK, response)

}

func (h *movieHandler) AddMovie(c echo.Context) error {
	movie := new(service.RequestMovie)
	if err := c.Bind(movie); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	err := h.movieService.CheckExistsMovie(*movie)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newMovie, _ := h.movieService.AddMovie(*movie)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	auth_token, err := h.authService.GetAccessToken(newMovie.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.MovieResponseFormatter(newMovie, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "movie successfully added", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) GetAllMovies(c echo.Context) error {
	movies, err := h.movieService.GetAllMovies()
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	// auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	var finalUserData []service.ResponseMovie
	for _, movie := range movies {
		auth_token, err := h.authService.GetAccessToken(movie.ID)
		if err != nil {
			response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

			return c.JSON(http.StatusInternalServerError, response)
		}
		userData := service.MovieResponseFormatter(movie, auth_token)
		finalUserData = append(finalUserData, userData)
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "get all movies succeeded", finalUserData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) AddGenreMovie(c echo.Context) error {
	genreMovie := new(service.RequestGenreMovie)
	if err := c.Bind(genreMovie); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	err := h.movieService.CheckExistsGenreMovie(*genreMovie)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newGenreMovie, _ := h.movieService.AddGenreMovie(*genreMovie)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	auth_token, err := h.authService.GetAccessToken(newGenreMovie.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.GenreMovieResponseFormatter(newGenreMovie, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "add genre movie succeeded", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) AddMovieReview(c echo.Context) error {
	movieReview := new(service.RequestMovieReview)
	if err := c.Bind(movieReview); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	newMovieReview, err := h.movieService.AddMovieReview(*movieReview)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	auth_token, err := h.authService.GetAccessToken(newMovieReview.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.MovieReviewResponseFormatter(newMovieReview, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "movie review successfully added", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) GetMoviewReview(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	movieReview, err := h.movieService.GetMovieReview(uint(id))
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	// auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	auth_token, err := h.authService.GetAccessToken(movieReview.ID)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	userData := service.MovieReviewResponseFormatter(*movieReview, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "get movie review successfull", userData)

	return c.JSON(http.StatusOK, response)
}
