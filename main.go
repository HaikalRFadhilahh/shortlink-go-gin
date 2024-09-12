package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/config"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/controllers"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/middleware"
	"github.com/HaikalRFadhilahh/shortlink-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Read .env File
	godotenv.Load()

	// Get Default .env
	APP_MODE := helper.GetEnv("APP_MODE", "production")
	HOST := helper.GetEnv("HOST", "0.0.0.0")
	PORT := helper.GetEnv("PORT", "3000")
	serveString := fmt.Sprintf("%s:%s", HOST, PORT)

	if APP_MODE != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Connection Database
	db, err := config.CreateConnection()
	if err == nil {
		log.Print("Database Connected!")
	}

	// Migrations
	db.AutoMigrate(&models.User{}, &models.Link{})

	// Controllers
	userController := &controllers.UserController{
		DB: db,
	}
	linkController := &controllers.LinkController{
		DB: db,
	}

	// Declare Gin Framework
	r := gin.Default()

	// All Middleware
	r.Use(middleware.ErrorMiddleware())

	// Route
	// Primary Link for redirect
	r.GET("/:alias", linkController.GetRedirectLink)

	// User Route
	r.POST("/auth/register", userController.Register)
	r.POST("/auth/login", userController.Login)

	// Link Route
	linkRoutesGroup := r.Group("/link")
	linkRoutesGroup.Use(middleware.CheckAuth())
	{
		linkRoutesGroup.POST("/create", linkController.CreateLink)
		linkRoutesGroup.DELETE("/delete/:idLink", linkController.DeleteLink)
		linkRoutesGroup.GET("/getall", linkController.GetAllLink)
	}
	// Not Found Route
	r.Use(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, helper.ErrorReponse{
			StatusCode: http.StatusNotFound,
			Status:     "error",
			Message:    "Not Found!",
		})
	})

	// Listen and Serve Server
	err = r.Run(serveString)
	if err != nil {
		log.Fatal(err.Error())
	}
}
