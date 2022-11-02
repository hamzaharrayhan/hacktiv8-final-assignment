package controller

import (
	"final-assignment/helper"
	"final-assignment/helper/input"
	"final-assignment/helper/response"
	"final-assignment/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var socialMediaService service.SocialMediaService

func NewSocialMediaController(socialMedia service.SocialMediaService, user service.UserService) {
	socialMediaService = socialMedia
	userService = user
}

func AddNewSocialMedia(c *gin.Context) {
	var input input.SocialMediaInput

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newSocialMedia, err := socialMediaService.CreateSocialMedia(input, currentUser)

	if err != nil {
		response := helper.JSONResponse("failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSocialMediaResponse := response.SocialMediaCreateResponse{
		ID:        newSocialMedia.ID,
		Name:      newSocialMedia.Name,
		URL:       newSocialMedia.URL,
		UserID:    newSocialMedia.UserID,
		CreatedAt: newSocialMedia.CreatedAt,
	}

	response := helper.JSONResponse("created", newSocialMediaResponse)
	c.JSON(http.StatusOK, response)
}

func DeleteSocialmedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	socmedID := c.Param("socialMediaId")

	idSocmed, err := strconv.Atoi(socmedID)
	if err != nil {
		response := helper.JSONResponse("failed", "social media not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if idSocmed == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	socialMedia, err := socialMediaService.GetSocialMediaByID(idSocmed)
	if err != nil {
		response := helper.JSONResponse("failed", "social media not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if currentUser != socialMedia.UserID {
		response := helper.JSONResponse("failed", "cannot delete social media that not belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = socialMediaService.DeleteSocialMedia(idSocmed)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("ok", "success deleted social media!")
	c.JSON(http.StatusOK, response)
}

func GetSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	socialmedia, err := socialMediaService.GetSocialMedia(currentUser)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	user, err := userService.GetUserByID(currentUser)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("ok", response.GetAllSocialMedia(socialmedia, user))
	c.JSON(http.StatusOK, response)
}

func UpdateSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	update := input.SocialMediaInput{}
	log.Println("updatenya apaaaaaaaaaaaaaa", update)
	err := c.ShouldBindJSON(&update)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	socmedID := c.Param("socialMediaId")

	idSocmed, err := strconv.Atoi(socmedID)
	if err != nil {
		response := helper.JSONResponse("failed", "social media id not found")
		c.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	if idSocmed == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	socialMedia, err := socialMediaService.GetSocialMediaByID(idSocmed)
	if err != nil {
		response := helper.JSONResponse("failed", "social media not found")
		c.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	if currentUser != socialMedia.UserID {
		response := helper.JSONResponse("failed", "cannot delete social media that not belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	queryResult, err := socialMediaService.UpdateSocialMedia(idSocmed, update)

	if queryResult.ID == 0 {
		response := helper.JSONResponse("failed", "update social media failed")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	updated, err := socialMediaService.GetSocialMediaByID(idSocmed)
	log.Println("updatedddddd", updated)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("Your social media has been successfully deleted", updated)
	c.JSON(http.StatusOK, response)
}
