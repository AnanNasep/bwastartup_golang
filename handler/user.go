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

func (h *userHandler) Login(c *gin.Context){
	//user memasukan input (email & password)
	//input ditangkap oleh handler
	//mapping dari input user ke input struct
	//input struct passing service
	//di service dengan bantuan repository user dengan email x
	//mencocokan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}		

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return		
	}
	// klo gk error lanjut...
	loggedinUser, err := h.userService.Login(input)

	if err != nil{
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return	
			
	}

	formatter := user.FormatUser(loggedinUser, "tokentoknetoknetoken")

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email di-mapping ke struct input
	// struct input di-pasing ke service
	// service akan memanggil - email sudah ada atau belum
	// repository - db 
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		//pemanggilan dari helper
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil{
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)	
		c.JSON(http.StatusUnprocessableEntity, response)
		return		
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}
	var metaMessage string

	if isEmailAvailable{
		metaMessage = "Email is available"
	}else{
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)	
	c.JSON(http.StatusOK, response)
}