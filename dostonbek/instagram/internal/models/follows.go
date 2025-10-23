package models

import "time"

type (
	Follow struct {
		ID           int       `json:"id"`
		FollowerId   int       `json:"follower_id"`
		FollowingId int       `json:"following_id"`
		CreatedAt    time.Time `json:"created_at"`
	}
)
