package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/middleware"
)

type movieHandler struct {
	userService  service.UserServices
	movieService service.MovieServices
	authService  auth.AuthService
}

func NewMovieHandler(userService service.UserServices, movieService service.MovieServices, authService auth.AuthService) *movieHandler {
	return &movieHandler{userService, movieService, authService}
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
	userData := service.GenreResponseFormatter(newGenre)
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
	// if err != nil {
	// 	response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

	// 	return c.JSON(http.StatusInternalServerError, response)
	// }

	var finalUserData []service.ResponseGenre
	for _, genre := range genres {
		userData := service.GenreResponseFormatter(genre)
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
	genres, _ := h.movieService.GetAllGenresByMovieID(newMovie.ID)
	userData := service.MovieResponseFormatter(newMovie, genres)
	response := helper.ResponseFormatter(http.StatusOK, "success", "movie successfully added", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) GetAllMovies(c echo.Context) error {
	movies, err := h.movieService.GetAllMovies()
	fmt.Printf("\n movieHandler GetAllMovies: %+v \n", movies)

	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	// auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	// if err != nil {
	// 	response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

	// 	return c.JSON(http.StatusInternalServerError, response)
	// }
	var finalUserData []service.ResponseMovie
	for _, movie := range movies {
		genres, _ := h.movieService.GetAllGenresByMovieID(movie.ID)
		userData := service.MovieResponseFormatter(movie, genres)
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

	userData := service.GenreMovieResponseFormatter(newGenreMovie)
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
	user, _ := h.userService.GetUserByID(newMovieReview.UserID)
	movie, _ := h.movieService.GetMovieByID(newMovieReview.MovieID)
	userData := service.MovieReviewResponseFormatter(newMovieReview, *user, *movie)
	response := helper.ResponseFormatter(http.StatusOK, "success", "movie review successfully added", userData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) GetMoviewReview(c echo.Context) error {
	movieID, _ := strconv.Atoi(c.Param("id"))
	movieReviews, err := h.movieService.GetMovieReview(uint(movieID))
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}
	// auth_token, err := h.authService.GetAccessToken(newGenre.ID)
	// if err != nil {
	// 	response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

	// 	return c.JSON(http.StatusInternalServerError, response)
	// }
	var finalData []service.ResponseMovieReview
	for _, movieReview := range movieReviews {
		user, _ := h.userService.GetUserByID(movieReview.UserID)
		movie, _ := h.movieService.GetMovieByID(movieReview.MovieID)
		moviewReviewData := service.MovieReviewResponseFormatter(movieReview, *user, *movie)
		finalData = append(finalData, moviewReviewData)
	}
	response := helper.ResponseFormatter(http.StatusOK, "success", "get movie review successfull", finalData)

	return c.JSON(http.StatusOK, response)
}

func (h *movieHandler) AdminAllowedAccess(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	role := claims.Role
	id := claims.ID
	if role != "admin" {
		return c.String(http.StatusForbidden, fmt.Sprintf("\n Access not allowed for this id: %v role:%s!\n", id, role))
	}
	return c.String(http.StatusOK, fmt.Sprintf("\n Welcome %s! \n", role))

}
