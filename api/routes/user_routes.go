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

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	log.Println("INSIDE UserRoutes.Route")
	db := database.GetDBInstance()
	repository.InitDBTable(db)
	repository.DBSeed(db)
	// db.AutoMigrate(User{}, Role{}, Permission{}, RolePermission{}, movie.Genre{}, movie.Movie{}, movie.GenreMovie{}, movie.MovieReview{})
	repo := repository.NewRepository(db)
	userService := service.NewUserServices(repo)
	authService := auth.NewAuthService()
	userHandler := handler.NewUserHandler(userService, authService)

	return []helper.Route{
		{
			Method: echo.POST,
			Path:   "/register",
			// Handler: userHandler.UserRegistration,
			Handler: userHandler.UserRegistrationSendEmail,
		},
		{
			Method:  echo.GET,
			Path:    "/register/confirmation",
			Handler: userHandler.UserRegisterConfirmation,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: userHandler.UserLogin,
		},
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
		},
		{
			Method:  echo.GET,
			Path:    "/user",
			Handler: userHandler.GetUser,
		},
		{
			Method:  echo.PUT,
			Path:    "/user",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:     echo.GET,
			Path:       "/secret",
			Handler:    userHandler.SecretResource,
			Middleware: []echo.MiddlewareFunc{middleware.JwtMiddleWare()},
		},
	}
}
