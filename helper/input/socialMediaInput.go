package input

type SocialMediaInput struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"social_media_url"`
}
