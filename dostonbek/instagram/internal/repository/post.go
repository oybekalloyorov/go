package repository

import (
	"database/sql"
	"log"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"time"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(dbconn *sql.DB) *PostRepo {
	return &PostRepo{
		db: dbconn,
	}
}

func (p *PostRepo) CreatePost(obj *models.Post) (*models.Post, error) {
	obj.CreatedAt = time.Now()

	query := `
		INSERT INTO post (title,description, likes_count, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, description, likes_count, created_at;
	`

	var result models.Post

	res := p.db.QueryRow(query, obj.Title, obj.Description, obj.LikesCount, obj.CreatedAt)

	if err := res.Scan(&result.ID, &result.Title, &result.Description, &result.LikesCount, &result.CreatedAt); err != nil {
		log.Println("Failed to parse")
		return nil, err
	}

	return &result, nil
}

func (p *PostRepo) GetPostByID(id int) (*models.Post, error) {
	query := `
		SELECT id, title, description, likes_count, created_at from post where id=$1;
	`
	var object models.Post
	err := p.db.QueryRow(query, id).Scan(&object.ID, &object.Title, &object.Description,
		&object.LikesCount, &object.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (p *PostRepo) GetAllPosts() ([]*models.Post, error) {
	query := `
		SELECT id, title, description, likes_count, created_at from post;
	`
	var result []*models.Post
	res, err := p.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer res.Close()

	for res.Next() {
		var obj models.Post
		if err := res.Scan(&obj.ID, &obj.Title, &obj.Description, &obj.LikesCount, &obj.CreatedAt); err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}

		result = append(result, &obj)
		if err := res.Err(); err != nil {
			log.Println(err.Error())
		}
	}
	return result, nil
}

func (p *PostRepo) UpdatePost(obj *models.Post) (*models.Post, error) {
	query := `
        UPDATE post
        SET
            title = COALESCE($2, title),
            description = COALESCE($3, description),
            likes_count = COALESCE($4, likes_count)
        WHERE id = $1
        RETURNING id, description, title, likes_count, created_at;
    `
	var result models.Post

	row := p.db.QueryRow(
		query,
		obj.ID,
		obj.Title,
		obj.Description,
		obj.LikesCount,
	)
	err := row.Scan(
		&result.ID,
		&result.Description,
		&result.Title,
		&result.LikesCount,
		&result.CreatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &result, nil
}

func (p *PostRepo) DeletePostByID(id int) error {
	query := `DELETE from post where id=$1;`

	res, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
