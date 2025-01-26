package auth_client

type CreateUserResponse struct {
	ID int `json:"id" validate:"required"`
}

type TokensResponse struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
