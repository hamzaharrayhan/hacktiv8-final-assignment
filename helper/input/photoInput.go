package input

type PhotoInput struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type UpdatePhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}
