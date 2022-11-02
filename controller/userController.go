package controller

import (
	"final-assignment/helper"
	"final-assignment/helper/input"
	"final-assignment/helper/response"
	"final-assignment/middleware"
	"final-assignment/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var userService service.UserService

func NewUserController(service service.UserService) {
	userService = service
}

func RegisterUser(c *gin.Context) {
	var userInput input.UserInput

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errorMessages := gin.H{
			"errors": err,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailValid := userService.EmailValidator(userInput.Email)
	if !isEmailValid {
		response := helper.JSONResponse("failed", gin.H{"error": "Invalid email format"})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if len(userInput.Password) <= 6 {
		response := helper.JSONResponse("failed", gin.H{"error": "Password length must be at least 6 characters"})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if userInput.Age <= 8 {
		response := helper.JSONResponse("failed", gin.H{"error": "You must be older or equal to 8 to register an account"})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := userService.CreateUser(userInput)
	if err != nil {
		errorMessages := gin.H{
			"errors": err,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userMap := response.UserResponse{
		Age:      newUser.Age,
		Email:    newUser.Email,
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	response := helper.JSONResponse("success", userMap)
	c.JSON(http.StatusCreated, response)
}

func LoginUser(c *gin.Context) {
	var loginInput input.UserLoginInput

	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := userService.GetUserByEmail(loginInput.Email)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if user.ID == 0 {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil {
		response := helper.JSONResponse("failed", "password not match")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	jwtService := middleware.NewService()
	token, err := jwtService.GenerateToken(user.ID)
	if err != nil {
		response := helper.JSONResponse("failed", "login failed")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.JSONResponse("logged in", gin.H{"token": token})
	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
	var userInput input.UserUpdateInput

	currentUser := c.MustGet("currentUser").(int)

	fmt.Println(currentUser)
	userID, err := strconv.Atoi(c.Param("userId"))
	fmt.Println(userID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = c.ShouldBindJSON(&userInput)
	fmt.Println(userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailValid := userService.EmailValidator(userInput.Email)
	if !isEmailValid {
		response := helper.JSONResponse("failed", gin.H{"error": "Invalid email format"})
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = userService.UpdateUser(userID, userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedUser, err := userService.GetUserByID(currentUser)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.JSONResponse("user updated", updatedUser)
	c.JSON(http.StatusOK, response)
}

func DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)
	user, err := userService.GetUserByID(currentUser)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = userService.DeleteUser(user.ID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.JSONResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.JSONResponse("success", gin.H{"message": "Your account has been successfully deleted"})
	c.JSON(http.StatusOK, response)
}
