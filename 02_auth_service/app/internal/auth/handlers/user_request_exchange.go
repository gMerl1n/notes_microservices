package handlers

type CreateUserRequest struct {
	Name           string
	Surname        string
	Age            int
	Email          string
	Password       string
	RepeatPassword string
}

type LoginUserRequest struct {
	Email    string
	Password string
}

type RefreshTokensRequest struct {
	Token string
}
