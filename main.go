package main

import (
	"backendEkost/auth"
	"backendEkost/handler"
	"backendEkost/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/ekostdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	//Test Validasi Token
	// // token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo2fQ.ErbRHHI-DYqCEwjmRfuBa60a40Slygl7jnXYi0Uq3bg")
	// // if err != nil {
	// // 	fmt.Println("EROOR")
	// // 	fmt.Println("EROOR")
	// // 	fmt.Println("EROOR")
	// // }
	// // if token.Valid {
	// // 	fmt.Println("VALID")
	// // 	fmt.Println("VALID")
	// // 	fmt.Println("VALID")
	// // } else {
	// // 	fmt.Println("INVALID")
	// // 	fmt.Println("INVALID")
	// // 	fmt.Println("INVALID")
	// // }
	// fmt.Println(authService.GenerateToken(1001))

	//====================================================
	// Test Service
	// input := user.LoginInput{
	// 	Email:    "christiantolie99@gmail.com",
	// 	Password: "yaudalah",
	// }
	// user, err := userService.Login(input)
	// if err != nil {
	// 	fmt.Println("Terjadi Kesalahan")
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(user.Email)
	// fmt.Println(user.Name)
	//=======================================================
	// userByEmail, err := userRepository.FindByEmail("Christiantolie99@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1") //API Versioning

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "test simpan dari service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "anak band"
	// userInput.Password = "ini password"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "test simpen",
	// }

	// userRepository.Save(user)
	// fmt.Println("Database Connected")

	// var users []user.User
	// length := len(users)

	// fmt.Println(length)

	// db.Find(&users)

	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("========")
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/ekostdb?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
