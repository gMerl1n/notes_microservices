package models

type NoteCreateRequest struct {
	UserID       int    `json:"user_id" validate:"required"`
	CategoryName string `json:"category_name" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Body         string `json:"body" validate:"required"`
}

type NoteGetRequestByID struct {
	NoteID int `json:"note_id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}

type NotesGetRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

type NoteRemoveRequestByID struct {
	NoteID int `json:"note_id" validate:"required"`
	UserID int `json:"user_id" validate:"required"`
}

type NotesRemove struct {
	UserID int `json:"user_id" validate:"required"`
}
