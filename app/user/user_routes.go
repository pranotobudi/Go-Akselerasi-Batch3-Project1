package user

import (
	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/database"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/middleware"
)

type UserRoutes struct{}

func (r UserRoutes) Route() []helper.Route {
	db := database.GetDBInstance()
	db.AutoMigrate(User{})
	userRepo := NewRepository(db)
	userService := NewServices(userRepo)
	authService := auth.NewAuthService()
	userHandler := NewHandler(userService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/register",
			Handler: userHandler.UserRegistration,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: userHandler.UserLogin,
		},
		{
			Method:     echo.GET,
			Path:       "/secret",
			Handler:    userHandler.SecretResource,
			Middleware: []echo.MiddlewareFunc{middleware.JwtMiddleWare()},
		},
	}
}
