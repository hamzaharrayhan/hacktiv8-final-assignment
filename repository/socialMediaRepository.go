package repository

import (
	"final-assignment/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Save(socialmedia model.SocialMedia) (model.SocialMedia, error)
	FindByUserID(ID int) ([]model.SocialMedia, error)
	Update(socialmedia model.SocialMedia, ID int) (model.SocialMedia, error)
	Delete(ID int) (model.SocialMedia, error)
	FindByID(ID int) (model.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Save(socialmedia model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Create(&socialmedia).Error

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialMediaRepository) FindByID(ID int) (model.SocialMedia, error) {
	socialmedia := model.SocialMedia{}

	err := r.db.Where("id = ?", ID).Find(&socialmedia).Error

	if err != nil {
		return socialmedia, err
	}

	return socialmedia, nil
}

func (r *socialMediaRepository) FindByUserID(ID int) ([]model.SocialMedia, error) {
	var socialmedia []model.SocialMedia

	err := r.db.Where("user_id = ?", ID).Find(&socialmedia).Error

	if err != nil {
		return []model.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialMediaRepository) Update(socialmedia model.SocialMedia, ID int) (model.SocialMedia, error) {
	err := r.db.Where("id = ?", ID).Updates(&socialmedia).Error

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (r *socialMediaRepository) Delete(ID int) (model.SocialMedia, error) {
	socialDeleted := model.SocialMedia{}

	err := r.db.Where("id = ?", ID).Delete(&socialDeleted).Error

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialDeleted, err
}
