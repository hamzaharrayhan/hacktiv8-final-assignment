package repository

import (
	"final-assignment/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
	Update(user model.User) (model.User, error)
	Delete(ID int) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindByID(ID int) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user model.User) (model.User, error) {
	fmt.Println(user)
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user model.User) (model.User, error) {
	err := r.db.Where("id = ?", user.ID).Updates(&user).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ID int) (model.User, error) {
	userDeleted := model.User{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&userDeleted).Error

	if err != nil {
		return model.User{}, err
	}

	return userDeleted, err
}

func (r *userRepository) FindByEmail(email string) (model.User, error) {
	user := model.User{}

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return model.User{}, nil
	}

	return user, nil
}

func (r *userRepository) FindByID(ID int) (model.User, error) {
	user := model.User{}

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
