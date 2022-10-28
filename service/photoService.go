package service

import (
	"final-assignment/helper/input"
	"final-assignment/model"
	"final-assignment/repository"
)

type PhotoService interface {
	CreatePhoto(photoInput input.PhotoInput, idUser int) (model.Photo, error)
	DeletePhoto(ID int) (model.Photo, error)
	GetPhotosUser(idUser int) ([]model.Photo, error)
	GetPhotoByID(idPhoto int) (model.Photo, error)
	UpdatePhoto(ID int, input input.UpdatePhoto) (model.Photo, error)
}

type photoService struct {
	photoRepo repository.PhotoRepository
}

func NewPhotoService(photoRepo repository.PhotoRepository) *photoService {
	return &photoService{photoRepo}
}

func (s *photoService) CreatePhoto(input input.PhotoInput, idUser int) (model.Photo, error) {
	newPhoto := model.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   idUser,
	}

	createNewPhoto, err := s.photoRepo.Save(newPhoto)

	if err != nil {
		return model.Photo{}, err
	}

	return createNewPhoto, nil

}

func (s *photoService) GetPhotosUser(idUser int) ([]model.Photo, error) {
	photos, err := s.photoRepo.FindByUserID(idUser)

	if err != nil {
		return []model.Photo{}, err
	}

	return photos, nil
}

func (s *photoService) DeletePhoto(ID int) (model.Photo, error) {
	photoQuery, err := s.photoRepo.FindByID(ID)

	if err != nil {
		return model.Photo{}, err
	}

	if photoQuery.ID == 0 {
		return model.Photo{}, nil
	}

	photoDeleted, err := s.photoRepo.Delete(ID)

	if err != nil {
		return model.Photo{}, err
	}

	return photoDeleted, nil
}

func (s *photoService) GetPhotoByID(idPhoto int) (model.Photo, error) {
	photoQuery, err := s.photoRepo.FindByID(idPhoto)

	if err != nil {
		return model.Photo{}, err
	}

	if photoQuery.ID == 0 {
		return model.Photo{}, nil
	}

	return photoQuery, nil
}

func (s *photoService) UpdatePhoto(ID int, input input.UpdatePhoto) (model.Photo, error) {

	photoResult, err := s.photoRepo.FindByID(ID)

	if err != nil {
		return model.Photo{}, err
	}

	if photoResult.ID == 0 {
		return model.Photo{}, nil
	}

	updatedPhoto := model.Photo{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoURL: input.PhotoURL,
		UserID:   photoResult.UserID,
	}

	photoUpdate, err := s.photoRepo.Update(updatedPhoto, ID)

	if err != nil {
		return model.Photo{}, err
	}

	return photoUpdate, nil
}
