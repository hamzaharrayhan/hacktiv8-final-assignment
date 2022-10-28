package service

import (
	"final-assignment/helper/input"
	"final-assignment/model"
	"final-assignment/repository"
)

type CommentService interface {
	CreateComment(input input.CommentInput, idUser int) (model.Comment, error)
	GetComment(UserID int) ([]model.Comment, error)
	DeleteComment(ID int) (model.Comment, error)
	UpdateComment(ID int, input input.CommentUpdateInput) (model.Comment, error)
	GetCommentByID(commentID int) (model.Comment, error)
	GetCommentsByPhotoID(photoID int) ([]model.Comment, error)
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) *commentService {
	return &commentService{commentRepo}
}

func (s *commentService) CreateComment(input input.CommentInput, idUser int) (model.Comment, error) {
	newComment := model.Comment{
		Message: input.Message,
		PhotoID: input.PhotoID,
		UserID:  idUser,
	}

	createNewcomment, err := s.commentRepo.Save(newComment)

	if err != nil {
		return model.Comment{}, err
	}

	return createNewcomment, nil
}

func (s *commentService) GetComment(UserID int) ([]model.Comment, error) {
	comment, err := s.commentRepo.FindByUserID(UserID)
	if err != nil {
		return []model.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) DeleteComment(ID int) (model.Comment, error) {
	comment, err := s.commentRepo.FindByID(ID)

	if err != nil {
		return model.Comment{}, err
	}

	if comment.ID == 0 {
		return model.Comment{}, nil
	}

	Deleted, err := s.commentRepo.Delete(ID)

	if err != nil {
		return model.Comment{}, err
	}

	return Deleted, nil
}

func (s *commentService) UpdateComment(ID int, input input.CommentUpdateInput) (model.Comment, error) {

	Result, err := s.commentRepo.FindByID(ID)

	if err != nil {
		return model.Comment{}, err
	}

	if Result.ID == 0 {
		return model.Comment{}, nil
	}

	updated := model.Comment{
		Message: input.Message,
	}

	commentUpdate, err := s.commentRepo.Update(updated, ID)

	if err != nil {
		return model.Comment{}, err
	}

	return commentUpdate, nil
}

func (s *commentService) GetCommentByID(commentID int) (model.Comment, error) {
	comment, err := s.commentRepo.FindByID(commentID)
	if err != nil {
		return model.Comment{}, err
	}

	return comment, nil
}

func (s *commentService) GetCommentsByPhotoID(photoID int) ([]model.Comment, error) {
	comments, err := s.commentRepo.FindByPhotoID(photoID)

	if err != nil {
		return []model.Comment{}, err
	}

	return comments, nil
}
