package repository

import (
	"database/sql"
	"log"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

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

func (u *UserRepo) GetUserById(id int) (*models.UserModel, error) {
	query := `
		SELECT id, full_name, username, birth_of_year,bio, created_at from instagram_users where id=$1;
	`

	var user models.UserModel
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.FullName, &user.Username, &user.BirthOfYear, &user.Bio, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (u *UserRepo) GetAllUsers() ([]*models.UserModel, error) {
	query := `
		SELECT id, full_name, username, birth_of_year, bio, created_at from instagram_users 
	`
	var res []*models.UserModel
	rows, err := u.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var obj models.UserModel
		if err := rows.Scan(&obj.ID, &obj.FullName, &obj.Username, &obj.BirthOfYear, &obj.Bio, &obj.CreatedAt); err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}

		res = append(res, &obj)
		if err := rows.Err(); err != nil {
			log.Println(err.Error())
		}
	}

	return res, nil
}
func (u *UserRepo) UpdateUserByID(obj *models.UserModel) (*models.UserModel, error) {
	query := `
		UPDATE instagram_users
		SET
			full_name = COALESCE($2, full_name),
			username = COALESCE($3, username),
			birth_of_year = COALESCE($4, birth_of_year),
			bio = COALESCE($5, bio)
		WHERE id = $1
		RETURNING id, full_name, username, birth_of_year, bio, created_at;
	`

	var res models.UserModel
	row := u.db.QueryRow(
		query,
		obj.ID,
		obj.FullName,
		obj.Username,
		obj.BirthOfYear,
		obj.Bio,
	)

	err := row.Scan(
		&res.ID,
		&res.FullName,
		&res.Username,
		&res.BirthOfYear,
		&res.Bio,
		&res.CreatedAt,
	)

	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	return &res, err
}

/*
	GetByID bajarildi
	GetAllUsers bajarildi
	UpdateByID bajarildi
	DeleteByID
*/
