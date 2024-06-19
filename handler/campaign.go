package handler

import (
	"fmt"
	"kami-peduli/campaign"
	"kami-peduli/helper"
	"kami-peduli/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(serive campaign.Service) *campaignHandler {
	return &campaignHandler{serive}
}

// api/v1/campaigs
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// value Query adalah string, padahal user_id adalah INT. maka harus di convert ke string
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaign", http.StatusBadRequest, "error,", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("list of campaigns", http.StatusOK, "sucsess,", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign detail", http.StatusOK, "succsess", campaign.FormatterCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error,", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success to create campaign", http.StatusOK, "success,", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed update campaign", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("success to update campaign", http.StatusOK, "success,", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

// func (h *campaignHandler) UploadImage(c *gin.Context) {
// 	var input campaign.CreateCampaignImageInput

// 	err := c.ShouldBind(&input)
// 	if err != nil {
// 		errors := helper.FormatValidatationError(err)
// 		errorMessage := gin.H{"errors": errors}

// 		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	currentUser := c.MustGet("currentUser").(user.User)
// 	input.User = currentUser
// 	userID := currentUser.ID

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

// 	err = c.SaveUploadedFile(file, path)
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	_, err = h.service.SaveCampaignImage(input, path)
// 	if err != nil {
// 		data := gin.H{"is_uploaded": false}
// 		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

//		response := helper.APIResponse("Campaign image successfully uploaded", http.StatusOK, "success", gin.H{"is_uploaded": true})
//		c.JSON(http.StatusOK, response)
//	}
func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidatationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed update campaign", http.StatusUnprocessableEntity, "error,", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("succsess to upload campaign image", http.StatusOK, "error", data)
	c.JSON(http.StatusOK, response)
}
