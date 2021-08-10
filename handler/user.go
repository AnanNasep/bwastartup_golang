package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handler membutuhkan bantuan dari service
type userHandler struct {
	userService user.Service
	//jwt
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler{
	return &userHandler{userService, authService}
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
	
	//PAnggil JWT
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil{
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)		
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//manggil dari formater
	formatter := user.FormatUser(newUser, token)

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

	//PAnggil JWT
	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil{
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", nil)		
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)

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

func (h *userHandler) UploadAvatar(c *gin.Context){
	//Input dari user
	//simpan gambar di folder "images/"
	//di service kita panggil repository
	//JWT (sementara hardcode, seakan2 user yang udh login ID = 1)
	//repo ambil data user yang ID = 1
	//repo update data user simpan ke lokasi file
	//c.SaveUploadedFile(file, )
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}

	// dapat dari JWT
	currentUser := c.MustGet("CurrentUser").(user.User)
	userID := currentUser.ID

	//menentukan alamat penyimpanan gambar
	//"images/%d-%s" maksudnya adalah, ketika user meng-upload gambar, namanya jadi ("IDuser-namagambar")	
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return 	
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
	
}

func (h *userHandler) FetchUser(c *gin.Context){
	currentUser := c.MustGet("CurrentUser").(user.User)
	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Succesfully fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}