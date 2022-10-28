package repository

import (
	"final-assignment/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Save(photo model.Photo) (model.Photo, error)
	Delete(ID int) (model.Photo, error)
	FindByID(ID int) (model.Photo, error)
	FindByUserID(ID int) ([]model.Photo, error)
	Update(photo model.Photo, ID int) (model.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Save(photo model.Photo) (model.Photo, error) {
	err := r.db.Save(&photo).Error

	if err != nil {
		return model.Photo{}, err
	}

	return photo, nil
}

func (r *photoRepository) FindByID(ID int) (model.Photo, error) {
	photo := model.Photo{}

	err := r.db.Where("id = ?", ID).Find(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *photoRepository) FindByUserID(ID int) ([]model.Photo, error) {
	var photos []model.Photo

	err := r.db.Where("user_id = ?", ID).Find(&photos).Error

	if err != nil {
		return []model.Photo{}, err
	}

	return photos, nil
}

func (r *photoRepository) Delete(ID int) (model.Photo, error) {
	photoDeleted := model.Photo{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&photoDeleted).Error

	if err != nil {
		return model.Photo{}, err
	}

	return photoDeleted, err
}

func (r *photoRepository) Update(photo model.Photo, ID int) (model.Photo, error) {
	err := r.db.Where("id = ?", ID).Updates(&photo).Error

	if err != nil {
		return model.Photo{}, err
	}

	return photo, nil
}
