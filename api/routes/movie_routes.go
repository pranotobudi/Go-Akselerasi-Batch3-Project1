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
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/middleware"
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
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/genre", //PASS
			Handler: movieHandler.GetAllGenres,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "member"),
			},
		},
		{
			Method:  echo.POST,
			Path:    "/movie", //PASS
			Handler: movieHandler.AddMovie,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/movie", // PASS
			Handler: movieHandler.GetAllMovies,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "member"),
			},
		},
		{
			Method:  echo.POST,
			Path:    "/genre-movie", // PASS
			Handler: movieHandler.AddGenreMovie,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.POST,
			Path:    "/review", // PASS
			Handler: movieHandler.AddMovieReview,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("member"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/review/:id", // CHECK AGAIN
			Handler: movieHandler.GetMoviewReview,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin", "member"),
			},
		},
	}
}
