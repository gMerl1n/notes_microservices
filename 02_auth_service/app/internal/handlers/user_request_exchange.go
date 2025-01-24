package handlers

type CreateUserRequest struct {
	Name           string `json:"name" validate:"required"`
	Surname        string `json:"surname" validate:"required"`
	Age            int    `json:"age" validate:"required,gte=0"`
	Email          string `json:"email" validate:"required"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RefreshTokensRequest struct {
	Token string `json:"token" validate:"required"`
}
