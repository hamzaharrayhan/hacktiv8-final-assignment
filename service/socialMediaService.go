package service

import (
	"final-assignment/helper/input"
	"final-assignment/model"
	"final-assignment/repository"
)

type SocialMediaService interface {
	CreateSocialMedia(input input.SocialMediaInput, idUser int) (model.SocialMedia, error)
	DeleteSocialMedia(ID int) (model.SocialMedia, error)
	UpdateSocialMedia(ID int, input input.SocialMediaInput) (model.SocialMedia, error)
	GetSocialMedia(UserID int) ([]model.SocialMedia, error)
	GetSocialMediaByID(ID int) (model.SocialMedia, error)
}

type socialMediaService struct {
	socialMediaRepo repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{socialMediaRepo}
}

func (s *socialMediaService) CreateSocialMedia(input input.SocialMediaInput, idUser int) (model.SocialMedia, error) {
	newSocialMedia := model.SocialMedia{
		Name:   input.Name,
		URL:    input.URL,
		UserID: idUser,
	}

	createdSocialmedia, err := s.socialMediaRepo.Save(newSocialMedia)

	if err != nil {
		return model.SocialMedia{}, err
	}

	return createdSocialmedia, nil

}

func (s *socialMediaService) GetSocialMedia(UserID int) ([]model.SocialMedia, error) {
	socialmedia, err := s.socialMediaRepo.FindByUserID(UserID)

	if err != nil {
		return []model.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *socialMediaService) GetSocialMediaByID(ID int) (model.SocialMedia, error) {
	socialmedia, err := s.socialMediaRepo.FindByID(ID)

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialmedia, nil
}

func (s *socialMediaService) DeleteSocialMedia(ID int) (model.SocialMedia, error) {
	socialmedia, err := s.socialMediaRepo.FindByID(ID)

	if err != nil {
		return model.SocialMedia{}, err
	}

	if socialmedia.ID == 0 {
		return model.SocialMedia{}, nil
	}

	socialmediaDeleted, err := s.socialMediaRepo.Delete(ID)

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialmediaDeleted, nil
}

func (s *socialMediaService) UpdateSocialMedia(ID int, input input.SocialMediaInput) (model.SocialMedia, error) {

	result, err := s.socialMediaRepo.FindByID(ID)

	if err != nil {
		return model.SocialMedia{}, err
	}

	if result.ID == 0 {
		return model.SocialMedia{}, nil
	}

	result.Name = input.Name
	result.URL = input.URL

	socialmediaUpdate, err := s.socialMediaRepo.Update(result, ID)

	if err != nil {
		return model.SocialMedia{}, err
	}

	return socialmediaUpdate, nil
}
