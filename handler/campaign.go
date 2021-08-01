package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
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



//update campaign
func (h *campaignHandler)  UpdatedCampaign(c *gin.Context){
//user masukan input
//dikirim ke handler
//mappting input ada 2 (input dari user, dan juga input dari uri)
//pssing ke service (find campaign by id, tangkap parameter)
//repository update data campaign	
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil{
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	
	//cek user sekarang yg sedang login)
	currentUser := c.MustGet("CurrentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil{
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)	
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))	
	c.JSON(http.StatusOK, response)
}

//UPLOAD CAMPAIGN IMAGE
func (h *campaignHandler) UploadImage(c *gin.Context){
//handler tangkap input ubah ke struck input
//save image campaign ke suatu folder
//service (kondisi manggil point 2 di repository kemudian panggill repo point 1)
//repository :
//	1. create / save image ke tabel campaign_images
//	2. ubah is_primary true ke false

	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload campaign image", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//cek user sekarang 9yg sedang login)
	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser
	userID := currentUser.ID

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}

	//menentukan alamat penyimpanan gambar
	//"images/%d-%s" maksudnya adalah, ketika user meng-upload gambar, namanya jadi ("IDuser-namagambar")	
	path := fmt.Sprintf("campaignImages/%d-%s", userID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}	
	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign image successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)	
}