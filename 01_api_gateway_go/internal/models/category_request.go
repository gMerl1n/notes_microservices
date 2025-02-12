package models

type CategoryCreateRequest struct {
	UserID       int    `json:"user_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
}

type CategoryGetRequestByID struct {
	UserID     int `json:"user_id" validate:"required"`
	CategoryID int `json:"category_id" validate:"required"`
}

type CategoriesGetRequest struct {
	UserID     int `json:"user_id" validate:"required"`
	CategoryID int `json:"category_id" validate:"required"`
}

type CategoryRemoveRequestByID struct {
	UserID     int `json:"user_id" validate:"required"`
	CategoryID int `json:"category_id" validate:"required"`
}
