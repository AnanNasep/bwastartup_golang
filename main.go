package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
  
func main() {	
	//koneksi ke db
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}

	//JWT 
	authService := auth.NewService()
	//userRepository ini memanggil / passing dari/ke repository
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)
	// campaign
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	//transaction
	transactionRepository := transaction.NewRepository(db)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	
	router := gin.Default()
	//cors go get github.com/gin-contrib/cors
	router.Use(cors.Default())
	//buat akses gambar secara langsung
	router.Static("/images", "./images")
	//untuk gruping /api/v1
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	//authMiddleware  authMiddleware(authService, userService)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	//fetch user
	api.POST("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	//ambil campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	//ambil campaign detail
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	// input create campaign
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	//update campaign
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdatedCampaign)
	//upload campaign image
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
	//transaction
	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	//list transaction
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	//input create transaction
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	//ambil data notifikasi dari midtrans
	api.POST("/transactions/transactions", transactionHandler.GetNotification)
	router.Run()
}
 
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc{
	return func (c *gin.Context){
		//MIDLEWARE UPDATE
			//Ambi nilai header authorization: Bearer tokentokentoken
			//Ambil header autorization ambil nilai tokennya saja
			//Validasi token
			//Ambil user_id
			//Ambil user dari DB berdasarkan user_id lewat service
			//Set context isinya user
			authHeader := c.GetHeader("Authorization")
			
			if !strings.Contains(authHeader, "Bearer"){
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			//split, ambil token nya doang by space
			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}
			//validasi token
			token, err := authService.ValidateToken(tokenString)
			if err != nil{
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid{
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			userID := int(claim["user_id"].(float64))

			user, err := userService.GetUserByID(userID)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			// memanggil user yang sedang login
			c.Set("CurrentUser", user)
	}

}	
	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes Simpan dari service2"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "tukang gorengan"
	// userInput.Password = "password"

	// userService.RegisterUser(userInput)


// 	// input
// 	// handler mapping input ke struct
// 	// service mapping ke struct User
// 	// Repository save struct User ke db
// 	// db


	// fmt.Println("")
	// fmt.Println("Koneksi Ke database berhasil")
	// fmt.Println("")
	// //Yang ada didalam package user    package.type
	// var users []user.User
	
	// //Memanggil tabel user db.find tabel user
	// db.Find(&users)
	
	// for _, user := range users{
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=====")
	// }

	// //router buat user akses HTTP blabla/ 
	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()



// //fungsi handler itu seperti sebuah controller
// func handler(c *gin.Context){
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup_golang?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil{
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)
// 	//buat output JSON
// 	c.JSON(http.StatusOK, users)



// }