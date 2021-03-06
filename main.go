package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/joho/godotenv"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/routes"
)

func main() {
	fmt.Println("bismillah")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("failed to load .env file")
	}
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Pre(middleware.RemoveTrailingSlash())
	routes.DefineApiRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))

}
