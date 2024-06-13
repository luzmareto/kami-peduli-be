package main

import (
	"fmt"
	"kami-peduli/auth"
	"kami-peduli/handler"
	"kami-peduli/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:lala1010@tcp(127.0.0.1:3306)/kami_peduli?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	autService := auth.NewService()

	fmt.Println(autService.GenerateToken(1001))

	userHandler := handler.NewUserHandler(userService, autService)

	// routing grouping
	router := gin.Default()
	api := router.Group("/api/v1")

	// register
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)

	router.Run()
}
