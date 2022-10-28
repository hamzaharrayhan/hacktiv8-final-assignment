package input

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id" binding:"required"`
}

type CommentUpdateInput struct {
	Message string `json:"message" binding:"required"`
}
