package controller

import (
	"final-assignment/helper"
	"final-assignment/helper/input"
	"final-assignment/helper/response"
	"final-assignment/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var photoService service.PhotoService
var commentService service.CommentService

func NewPhotoController(photo service.PhotoService, comment service.CommentService, user service.UserService) {
	photoService = photo
	commentService = comment
	userService = user
}

func AddNewPhoto(c *gin.Context) {
	var input input.PhotoInput

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
	newPhoto, err := photoService.CreatePhoto(input, currentUser)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPhotoResponse := response.PhotoResponse{
		ID:        newPhoto.ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoURL:  input.PhotoURL,
		UserID:    currentUser,
		CreatedAt: newPhoto.CreatedAt,
	}

	response := helper.JSONResponse("created", newPhotoResponse)
	c.JSON(http.StatusOK, response)
}

func DeletePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	photoID := c.Param("photoId")

	idPhoto, err := strconv.Atoi(photoID)
	if err != nil {
		response := helper.JSONResponse("failed", "photo not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if idPhoto == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photo, err := photoService.GetPhotoByID(idPhoto)
	if err != nil {
		response := helper.JSONResponse("failed", "photo not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if currentUser != photo.UserID {
		response := helper.JSONResponse("failed", "cannot delete photo that not belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = photoService.DeletePhoto(idPhoto)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("ok", "photo deleted")
	c.JSON(http.StatusOK, response)
}

func GetPhotos(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photos, err := photoService.GetPhotosUser(currentUser)

	var photoResponse []response.GetPhotoWithUserDetail

	for _, photo := range photos {
		user, _ := userService.GetUserByID(photo.UserID)
		commentTmp, _ := commentService.GetCommentsByPhotoID(photo.ID)

		photoResponseTmp := response.GetPhotoWithUserDetail{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			CreatedAt: photo.CreatedAt,
			Comments:  commentTmp,
			User: response.UserPhoto{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
		}

		photoResponse = append(photoResponse, photoResponseTmp)
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("ok", photoResponse)
	c.JSON(http.StatusOK, response)
}

func UpdatePhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	updatePhoto := input.UpdatePhoto{}

	err := c.ShouldBindJSON(&updatePhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photoID := c.Param("photoId")

	idPhoto, err := strconv.Atoi(photoID)
	if err != nil {
		response := helper.JSONResponse("failed", "comment not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if idPhoto == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	photo, err := photoService.GetPhotoByID(idPhoto)
	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	if photo.UserID != currentUser {
		response := helper.JSONResponse("failed", "cannot update photo that not belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = photoService.UpdatePhoto(idPhoto, updatePhoto)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	photoUpdated, _ := photoService.GetPhotoByID(idPhoto)

	response := helper.JSONResponse("Your photo has been successfully deleted", photoUpdated)
	c.JSON(http.StatusOK, response)
}
