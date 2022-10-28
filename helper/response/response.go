package response

import (
	"final-assignment/model"
	"time"
)

type UserResponse struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type PhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUserPhoto struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Caption   string          `json:"caption"`
	PhotoURL  string          `json:"photo_url"`
	CreatedAt time.Time       `json:"created_at"`
	Comments  []model.Comment `json:"comments"`
}

type GetPhotoWithUserDetail struct {
	ID        int             `json:"id"`
	Title     string          `json:"title"`
	Caption   string          `json:"caption"`
	PhotoURL  string          `json:"photo_url"`
	CreatedAt time.Time       `json:"created_at"`
	User      UserPhoto       `json:"user"`
	Comments  []model.Comment `json:"comments"`
}

type UserPhoto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
}

type GetCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	CreatedAt time.Time `json:"-"`
	Photo     model.Photo
}

func GetAllComment(comment model.Comment, photo model.Photo) GetCommentResponse {
	var response GetCommentResponse

	response.ID = comment.ID
	response.Message = comment.Message
	response.PhotoID = comment.PhotoID
	response.CreatedAt = comment.CreatedAt
	response.Photo = photo

	return response
}

type SocialMediaCreateResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"date"`
}

type GetSocialMedia struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"social_media_url"`
	UsedID    int       `json:"user_id"`
	CreatedAt time.Time `json:"date"`
	User      model.User
}

func GetAllSocialMedia(social []model.SocialMedia, user model.User) []GetSocialMedia {
	if len(social) == 0 {
		return []GetSocialMedia{}
	}

	var allSocialMedia []GetSocialMedia

	for _, socialmedia := range social {
		tmpSocialmedia := GetSocialMedia{
			ID:        socialmedia.ID,
			Name:      socialmedia.Name,
			URL:       socialmedia.URL,
			UsedID:    socialmedia.UserID,
			CreatedAt: socialmedia.CreatedAt,
			User: model.User{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
		}

		allSocialMedia = append(allSocialMedia, tmpSocialmedia)

	}

	return allSocialMedia
}
