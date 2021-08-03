package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HANDLER ADALAH SEMACAM CONTROLLER

//parameter di URI
//tangkap parameter mapping ke input struck
//panggil ke service, input struck sebagai parameter
//service, berbekal campaign_id kemudian bisa panggil repository
//repo mencari data transaction suatu campaign

type transactionHandler struct{
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler{
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context){
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil{
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error 1", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser
	
	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil{
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error 2", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Campaign's transactions", http.StatusOK,"success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}