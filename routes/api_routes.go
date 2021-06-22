package routes

import (
	"log"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/app/user"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handler{
		user.UserRoutes{},
	}

	var routes []helper.Route
	log.Println("WE'RE HERE handler: ", handlers)
	for _, handler := range handlers {
		// log.Println("WE'RE HERE routes: ", handler)
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/api/v1/movie_reviews")
	log.Println("WE'RE HERE route", routes)
	for _, route := range routes {
		log.Println("WE'RE HERsE: ", route)
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		}

	}
}
