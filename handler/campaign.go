package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// tangkap parameter di handler
// handler ke service
// service yang menentukan repository mana yang di call
// repository : FindAll, FindUserBtId
// akses ke db

type campaignHandler struct{
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler{
	return &campaignHandler{service}
}

//api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context){
	//user id di ubah ke string
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of campaigns", http.StatusBadRequest, "success", campaign.FormatCampaigns(campaigns))		
	c.JSON(http.StatusOK, response)

}