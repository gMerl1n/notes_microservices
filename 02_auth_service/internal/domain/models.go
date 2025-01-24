package domain

import "time"

type User struct {
	ID        int
	Name      string
	Surname   string
	Age       int
	Email     string
	Password  string
	RoleID    int
	UpdatedAt time.Time
	CreatedAt time.Time
}
