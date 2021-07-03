package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handler membutuhkan bantuan dari service
type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler{
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input dari user
	//map input dari user ke struct RegisterUserInput
	//struct di atas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		//pemanggilan dari helper
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		//manggil dari helper
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//input dari API 
	newUser, err := h.userService.RegisterUser(input)
	
	if err != nil{
		//manggil dari helper
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)		
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
// token, err := h.jwtService.GeneralToken()

	//manggil dari formater
	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	//manggil dari helper
	response := helper.APIResponse("Account has been registered", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)

}