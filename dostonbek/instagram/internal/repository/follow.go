package repository

import (
	"database/sql"
	"log"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"time"
)

type FollowRepo struct {
	db *sql.DB
}

func NewFollowRepo(dbconn *sql.DB) *FollowRepo{
	return &FollowRepo{
		db: dbconn,
	}
}

func (f *FollowRepo) CreateFollows(obj *models.Follow)(*models.Follow, error){
	obj.CreatedAt = time.Now()
	query := `
		INSERT INTO follows(follower_id, following_id, created_at)
		VALUES ($1,$2,$3)
		RETURNING id, follower_id, following_id, created_at;
	`
	row := f.db.QueryRow(query, obj.FollowerId, obj.FollowingId, obj.CreatedAt)

	var res models.Follow
	if err := row.Scan(&res.ID, &res.FollowerId, &res.FollowingId, &res.CreatedAt); err != nil {
		log.Println("Failed to parse")
		return nil, err
	}

	return &res, nil

}
