package handler

import (
	"kami-peduli/helper"
	"kami-peduli/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// logic handler
func (h *userHandler) RegisterUser(c *gin.Context) {
	// menangkap data input dari register
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Register account fail", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// token, err := h.jwtService.GenerateToken

	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	// output meta
	response := helper.APIResponse("account has been registerd", http.StatusOK, "succsess,", formatter)

	c.JSON(http.StatusOK, response)
}
