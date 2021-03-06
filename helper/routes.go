package helper

import "github.com/labstack/echo"

type Route struct {
	Method     string
	Path       string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

type Handler interface {
	Route() []Route
}
