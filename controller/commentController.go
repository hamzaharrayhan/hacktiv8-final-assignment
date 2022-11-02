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

func NewCommentController(comment service.CommentService, photo service.PhotoService) {
	commentService = comment
	photoService = photo
}

func AddNewComment(c *gin.Context) {
	var input input.CommentInput

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorsMessage := helper.FormatValidationError(err)

		response := helper.JSONResponse("failed", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newComment, err := commentService.CreateComment(input, currentUser)

	if err != nil {

		response := helper.JSONResponse("failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCommentResponse := response.CreateCommentResponse{
		ID:      newComment.ID,
		Message: newComment.Message,
		PhotoID: input.PhotoID,
		UserID:  currentUser,
	}

	response := helper.JSONResponse("created", newCommentResponse)
	c.JSON(http.StatusOK, response)
}

func DeleteComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	commentID := c.Param("commentId")
	idComment, err := strconv.Atoi(commentID)
	if err != nil {
		response := helper.JSONResponse("failed", "comment not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if idComment == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	comment, err := commentService.GetCommentByID(idComment)
	if err != nil {
		response := helper.JSONResponse("failed", "comment not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if comment.UserID != currentUser {
		response := helper.JSONResponse("failed", "you can only delete comments that belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = commentService.DeleteComment(idComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.JSONResponse("ok", "success deleted comment")
	c.JSON(http.StatusOK, response)
}

func GetComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	comments, _ := commentService.GetComment(currentUser)

	// query photo
	var allCommentsPhoto []response.GetCommentResponse
	for _, item := range comments {
		photo, _ := photoService.GetPhotoByID(item.PhotoID)
		allCommentsPhotoTmp := response.GetAllComment(item, photo)

		allCommentsPhoto = append(allCommentsPhoto, allCommentsPhotoTmp)
	}

	response := helper.JSONResponse("ok", allCommentsPhoto)
	c.JSON(http.StatusOK, response)
}

func UpdateComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.JSONResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	UpdateComment := input.CommentUpdateInput{}

	err := c.ShouldBindJSON(&UpdateComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	commentID := c.Param("commentId")
	idComment, err := strconv.Atoi(commentID)
	if err != nil {
		response := helper.JSONResponse("failed", "comment not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if idComment == 0 {
		response := helper.JSONResponse("failed", "id must be exist")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	comment, err := commentService.GetCommentByID(idComment)
	if err != nil {
		response := helper.JSONResponse("failed", "comment not found")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	if comment.UserID != currentUser {
		response := helper.JSONResponse("failed", "you can only delete comments that belongs to you")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = commentService.UpdateComment(idComment, UpdateComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.JSONResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	Updated, _ := commentService.GetCommentByID(idComment)

	if Updated.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, "comment not found")
		return
	}

	response := helper.JSONResponse("Your comment has been successfully deleted", Updated)
	c.JSON(http.StatusOK, response)
}
