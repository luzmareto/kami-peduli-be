package handler

import (
	"fmt"
	"kami-peduli/auth"
	"kami-peduli/helper"
	"kami-peduli/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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
		response := helper.APIResponse("Register account fail", http.StatusUnprocessableEntity, "error,", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account fail", http.StatusUnprocessableEntity, "error,", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error,", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// berhasil
	formatter := user.FormatUser(loggedinUser, token)

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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// field form menjadi avatar
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_upladed": false} //key meta
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	// lokasi penyimpanan avatar dan auto mengubah nama file yang di upload oleh user
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename) //otomatis akan memasukan id user ke db

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_upladed": false} //key meta
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_upladed": false} //key meta
		response := helper.APIResponse("failed upload avatar image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_upladed": true}
	response := helper.APIResponse("Avatar Succsessfully uploaded", http.StatusOK, "succsess", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")
	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
