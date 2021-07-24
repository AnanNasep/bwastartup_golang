package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

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

	//userRepository ini memanggil / passing dari/ke repository
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	//JWT 
	authService := auth.NewService()

	//userService.SaveAvatar(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()


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
}

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