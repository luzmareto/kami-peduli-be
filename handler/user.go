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

// input akan dikelola hanndler dan dipassing ke service untuk mencocokan email & pass
func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login fail", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// validasi agar ketika login gagal, maka prosesnya akan berhenti
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login fail", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// berhasil
	formatter := user.FormatUser(loggedinUser, "tokentokentokentoken")
	response := helper.APIResponse("Succsessfully Login", http.StatusOK, "succsess,", formatter)
	c.JSON(http.StatusOK, response)
}

// proses pengecekan email register
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking fail", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// gagal
	IsEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Server error"}
		response := helper.APIResponse("Email checking fail", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": IsEmailAvailable,
	}

	var metaMessage = "email has been registered"

	if IsEmailAvailable {
		metaMessage = "email is available"
	}

	// berhasil
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
