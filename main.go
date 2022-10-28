package main

import (
	"final-assignment/config"
	"final-assignment/controller"
	"final-assignment/middleware"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.StartDB()
	fmt.Println("Starting database")

	config.Setup()
	router := startServer()
	fmt.Println("Staring server")
	router.Run()
}

func startServer() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions, http.MethodPut},
		AllowHeaders:     []string{"Content-Type", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// user
	router.POST("/users/register", controller.RegisterUser)
	router.POST("/users/login", controller.LoginUser)
	router.PUT("/users/:userId", middleware.AuthMiddleware(), controller.UpdateUser)
	router.DELETE("/users", middleware.AuthMiddleware(), controller.DeleteUser)

	// photos
	router.POST("/photos", middleware.AuthMiddleware(), controller.AddNewPhoto)
	router.DELETE("/photos/:photoId", middleware.AuthMiddleware(), controller.DeletePhoto)
	router.GET("/photos", middleware.AuthMiddleware(), controller.GetPhotos)
	router.GET("/photos/:photoId", controller.GetPhoto)
	router.PUT("/photos/:photoId", middleware.AuthMiddleware(), controller.UpdatePhoto)

	// comments
	router.POST("/comments", middleware.AuthMiddleware(), controller.AddNewComment)
	router.DELETE("/comments/:commentId", middleware.AuthMiddleware(), controller.DeleteComment)
	router.GET("/comments", middleware.AuthMiddleware(), controller.GetComment)
	router.PUT("/comments/:commentId", middleware.AuthMiddleware(), controller.UpdateComment)

	// social media
	router.POST("/socialmedias", middleware.AuthMiddleware(), controller.AddNewSocialMedia)
	router.GET("/socialmedias", middleware.AuthMiddleware(), controller.GetSocialMedia)
	router.PUT("/socialmedias/:socialMediaId", middleware.AuthMiddleware(), controller.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", middleware.AuthMiddleware(), controller.DeleteSocialmedia)

	return router
}
