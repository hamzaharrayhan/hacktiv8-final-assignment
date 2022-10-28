package config

import (
	"final-assignment/controller"
	"final-assignment/repository"
	"final-assignment/service"
)

var userRepository repository.UserRepository
var userService service.UserService
var commentRepository repository.CommentRepository
var commentService service.CommentService
var socialMediaRepository repository.SocialMediaRepository
var socialMediaService service.SocialMediaService
var photoRepository repository.PhotoRepository
var photoService service.PhotoService

func Setup() {
	CreateRepositories()
	CreateServices()
	CreateController()
}

func CreateRepositories() {
	userRepository = repository.NewUserRepository(GetDB())
	photoRepository = repository.NewPhotoRepository(GetDB())
	commentRepository = repository.NewCommentRepository(GetDB())
	socialMediaRepository = repository.NewSocialMediaRepository(GetDB())
}

func CreateServices() {
	userService = service.NewUserService(userRepository)
	photoService = service.NewPhotoService(photoRepository)
	commentService = service.NewCommentService(commentRepository)
	socialMediaService = service.NewSocialMediaService(socialMediaRepository)
}

func CreateController() {
	controller.NewUserController(userService)
}
