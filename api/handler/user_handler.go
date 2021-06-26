package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project1/helper"
)

type userHandler struct {
	userService service.UserServices
	authService auth.AuthService
}

func NewUserHandler(userService service.UserServices, authService auth.AuthService) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) UserRegistration(c echo.Context) error {
	// user := &User{}
	user := new(service.RequestUser)
	if err := c.Bind(user); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	err := h.userService.CheckExistsEmail(*user)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	newUser, _ := h.userService.CreateUser(*user)
	if err != nil {
		errorFormatter := helper.ErrorFormatter(err)
		errorMessage := helper.M{"errors": errorFormatter}
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", errorMessage, nil)

		return c.JSON(http.StatusBadRequest, response)
	}

	role, _ := h.userService.GetRole(newUser.ID)
	fmt.Printf("\n Handler: userID: %v ROLE: %+v \n", newUser.ID, role)

	auth_token, err := h.authService.GetAccessToken(newUser.ID, role)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}
	userData := service.UserResponseFormatter(newUser, auth_token)
	response := helper.ResponseFormatter(http.StatusOK, "success", "user successfully registered", userData)

	return c.JSON(http.StatusOK, response)
}

func (u *userHandler) UserLogin(c echo.Context) error {
	userLogin := new(service.RequestUserLogin)
	if err := c.Bind(userLogin); err != nil {
		response := helper.ResponseFormatter(http.StatusBadRequest, "error", "invalid request", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	userAuth, err := u.userService.AuthUser(*userLogin)
	if err != nil {
		fmt.Println("We're IN HERE: USERLOGIN INSIDE")
		response := helper.ResponseFormatter(http.StatusUnauthorized, "error", err.Error(), nil)
		return c.JSON(http.StatusUnauthorized, response)
	}
	role, _ := u.userService.GetRole(userAuth.ID)

	auth_token, err := u.authService.GetAccessToken(userAuth.ID, role)
	if err != nil {
		response := helper.ResponseFormatter(http.StatusInternalServerError, "error", err.Error(), nil)

		return c.JSON(http.StatusInternalServerError, response)
	}

	userData := service.UserResponseFormatter(userAuth, auth_token)

	response := helper.ResponseFormatter(http.StatusOK, "success", "user authenticated", userData)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) SecretResource(c echo.Context) error {
	response := helper.M{"message": "this is secret route"}

	return c.JSON(http.StatusOK, response)
}
