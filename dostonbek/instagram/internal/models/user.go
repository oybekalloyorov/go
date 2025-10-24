package models

import "time"

type UserModel struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	Username    string    `json:"username"`
	BirthOfYear int       `json:"birth_of_year"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
}
