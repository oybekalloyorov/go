package models

import "time"

type (
	Post struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		LikesCount  int       `json:"likes_count"`
		CreatedAt   time.Time `json:"created_at"`
	}
)
