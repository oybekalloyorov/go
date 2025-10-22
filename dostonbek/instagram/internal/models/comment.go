package models

import "time"

type (
	Comment struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedBy int       `json:"created_by"`
		CreatedAt time.Time `json:"created_at"`
	}
)
