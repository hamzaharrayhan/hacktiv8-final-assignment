package repository

import (
	"final-assignment/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Save(comment model.Comment) (model.Comment, error)
	Delete(ID int) (model.Comment, error)
	FindByUserID(ID int) ([]model.Comment, error)
	Update(comment model.Comment, ID int) (model.Comment, error)
	FindByID(ID int) (model.Comment, error)
	FindByPhotoID(IDPhoto int) ([]model.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Save(comment model.Comment) (model.Comment, error) {
	err := r.db.Save(&comment).Error

	if err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) Delete(ID int) (model.Comment, error) {
	commentDeleted := model.Comment{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(&commentDeleted).Error

	if err != nil {
		return model.Comment{}, err
	}

	return commentDeleted, err
}

func (r *commentRepository) FindByUserID(ID int) ([]model.Comment, error) {
	var comment []model.Comment

	err := r.db.Where("user_id = ?", ID).Find(&comment).Error

	if err != nil {
		return []model.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) FindByID(ID int) (model.Comment, error) {
	comment := model.Comment{}

	err := r.db.Where("id = ?", ID).Find(&comment).Error

	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *commentRepository) Update(comment model.Comment, ID int) (model.Comment, error) {
	err := r.db.Where("id = ?", ID).Updates(&comment).Error

	if err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) FindByPhotoID(IDPhoto int) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("photo_id = ?", IDPhoto).Find(&comments).Error

	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}
