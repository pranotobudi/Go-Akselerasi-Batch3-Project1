package routes

import (
	"log"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/handler"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/repository"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/database"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
)

type MovieRoutes struct{}

func (r MovieRoutes) Route() []helper.Route {
	log.Println("INSIDE MovieRoutes.Route")
	db := database.GetDBInstance()
	repository.InitDBTable(db)
	repository.DBSeed(db)
	// db.AutoMigrate(User{}, Role{}, Permission{}, RolePermission{}, movie.Genre{}, movie.Movie{}, movie.GenreMovie{}, movie.MovieReview{})
	repo := repository.NewRepository(db)
	movieService := service.NewMovieServices(repo)
	authService := auth.NewAuthService()
	movieHandler := handler.NewMovieHandler(movieService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/genre", //PASS
			Handler: movieHandler.AddGenre,
		},
		{
			Method:  echo.GET,
			Path:    "/genre", //ONLY 1 ROW, NOT ALL, CHECK AGAIN
			Handler: movieHandler.GetAllGenres,
		},
		{
			Method:  echo.POST,
			Path:    "/movie", //PASS
			Handler: movieHandler.AddMovie,
		},
		{
			Method:  echo.GET,
			Path:    "/movie", // CAN'T ACCESS TABLE
			Handler: movieHandler.GetAllMovies,
		},
		{
			Method:  echo.POST,
			Path:    "/genre-movie", // PASS
			Handler: movieHandler.AddGenreMovie,
		},
		{
			Method:  echo.POST,
			Path:    "/review", // PASS
			Handler: movieHandler.AddMovieReview,
		},
		{
			Method:  echo.GET,
			Path:    "/review/:id", // CHECK AGAIN
			Handler: movieHandler.GetMoviewReview,
		},
	}
}
