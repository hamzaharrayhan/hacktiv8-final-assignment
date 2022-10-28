package service

import (
	"errors"
	"final-assignment/helper/input"
	"final-assignment/model"
	"final-assignment/repository"
	"log"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(input input.UserInput) (model.User, error)
	EmailValidator(email string) bool
	UpdateUser(id int, input input.UserUpdateInput) (model.User, error)
	GetUserByID(ID int) (model.User, error)
	DeleteUser(ID int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(input input.UserInput) (model.User, error) {
	user := model.User{}
	user.Username = input.Username
	user.Age = input.Age
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return model.User{}, err
	}

	user.Password = string(passwordHash)

	log.Println(user)
	newUser, err := s.userRepo.Register(user)
	if err != nil {
		return model.User{}, err
	}

	return newUser, nil
}
func (s *userService) UpdateUser(id int, input input.UserUpdateInput) (model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}

	if user.ID == 0 {
		return model.User{}, errors.New("user not found")
	}

	user.Email = input.Email
	user.Username = input.Username

	userUpdate, err := s.userRepo.Update(user)
	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *userService) DeleteUser(ID int) (model.User, error) {
	userQuery, err := s.userRepo.FindByID(ID)
	if err != nil {
		return model.User{}, err
	}

	if userQuery.ID == 0 {
		return model.User{}, nil
	}

	deletedUser, err := s.userRepo.Delete(ID)
	if err != nil {
		return model.User{}, err
	}

	return deletedUser, nil
}

func (s *userService) EmailValidator(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *userService) GetUserByID(ID int) (model.User, error) {
	user, err := s.userRepo.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return model.User{}, nil
	}

	return user, nil
}

func (s *userService) GetUserByEmail(email string) (model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}
