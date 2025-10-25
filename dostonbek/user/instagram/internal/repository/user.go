package repository

import (
	"database/sql"
	"instagram/internal/models"
	"log"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

/*
{
	"username": "prodonik",
	"full_name": "Dostonbek Soliyev",
	"bio":"backend developer"
}

{
	"username": "prodonik",
	"full_name": "Dostonbek Soliyev",
	"bio":"backend developer"
}
*/

func (u *UserRepo) CreateUser(req *models.UserModel) (*models.UserModel, error) {
	query := `
		INSERT INTO instagram_users(full_name, username, birth_of_year, bio)
		VALUES ($1, $2, $3, $4)
		RETURNING id, full_name, username, birth_of_year, bio, created_at;
	`

	row := u.db.QueryRow(query, req.FullName, req.Username, req.BirthOfYear, req.Bio)

	var response models.UserModel
	err := row.Scan(
		&response.ID,
		&response.FullName,
		&response.Username,
		&response.BirthOfYear,
		&response.Bio,
		&response.CreatedAt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &response, nil
}

/*
	GetByID
	GetAllUsers
	UpdateByID
	DeleteByID
*/
