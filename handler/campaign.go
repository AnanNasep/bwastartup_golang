package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
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

// CAMPAIGN DETAIL
func (h *campaignHandler) GetCampaign(c *gin.Context){
	//bentuk api ====>    api/v1/campaigns/1
	//handler	: mapping id yg di url ke struct input ke => service => call formater
	//service	: inputnya struct input => menangkap ID di url, manggi repository
	//repository : get campaign by id

	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil{
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaign_Detail, err := h.service.GetCampaignByID(input)
	if err != nil{
		response := helper.APIResponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}	
	response := helper.APIResponse("Campaign detail", http.StatusOK,"success", campaign.FormatCampaignDetail(campaign_Detail))
	c.JSON(http.StatusOK, response)
}

//CREATE CAMPAIGN
	//tangkap parameter dari user ke input struck 
	//ambil current user dari jwt/handler
	//panggil service, parameternya input struct (dan juga buat slug)
	//panggil repository untuk simpan data campaign baru

func (h *campaignHandler) CreateCampaign(c *gin.Context){
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil{
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)	
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))	
	c.JSON(http.StatusOK, response)
}