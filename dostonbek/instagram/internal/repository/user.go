package repository

import (
	"database/sql"
	"log"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo{
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(req *models.UserModel) (*models.UserModel, error){
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
	if err != nil{
		log.Println(err.Error())
		return nil, err
	}

	return &response, nil
}

func (u * UserRepo) GetUserById(id int) (*models.UserModel, error){
	query :=`
		SELECT id, full_name, username, birth_of_year,bio, created_at from instagram_users where id=$1;
	`

	var user models.UserModel
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.FullName, &user.Username, &user.BirthOfYear, &user.Bio, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil


}

/*
	GetByID bajarildi
	GetAllUsers
	UpdateByID
	DeleteByID
*/