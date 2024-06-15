package handler

import (
	"kami-peduli/campaign"
	"kami-peduli/helper"
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
	response := helper.APIResponse("list of campaigns", http.StatusOK, "sucsess,", campaigns)
	c.JSON(http.StatusOK, response)

}
