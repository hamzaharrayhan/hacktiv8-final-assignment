package input

type UserInput struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
