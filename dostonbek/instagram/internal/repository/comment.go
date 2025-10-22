package repository

import (
	"database/sql"
	"log"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"time"
)

type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(dbconn *sql.DB) *CommentRepo{
	return &CommentRepo{
		db: dbconn,
	}
}

func (c *CommentRepo) CreateComment(obj *models.Comment)(*models.Comment, error){
	obj.CreatedAt = time.Now()
	query := `
		INSERT INTO comments (title, content, created_by, created_at)
		VALUES ($1,$2, $3, $4)
		RETURNING id, title, content, created_by, created_at;
	`

	var result models.Comment

	res := c.db.QueryRow(query, obj.Title, obj.Content, obj.CreatedBy, obj.CreatedAt)

	if err := res.Scan(&result.ID, &result.Title, &result.Content, &result.CreatedBy, &result.CreatedAt); err != nil {
		log.Println("Failed to Println")
		return nil, err
	}

	return &result, nil
}
func (c *CommentRepo)GetCommentById(id int)(*models.Comment, error){
	query := `
		SELECT id, title, content, created_by, created_at from comments where id = $1
	`
	var object models.Comment
	err := c.db.QueryRow(query, id).Scan(
		&object.ID, 
		&object.Title, 
		&object.Content, 
		&object.CreatedBy, 
		&object.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &object, nil
}

func (c *CommentRepo) GetAllComments()([]*models.Comment, error){
	query := `
		SELECT id, title, content,created_by,created_at from comments
	`
	var result []*models.Comment
	rows, err := c.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){
		var obj models.Comment
		if err := rows.Scan(&obj.ID, &obj.Title, &obj.Content, &obj.CreatedBy, &obj.CreatedAt); err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}

		result = append(result, &obj)
		if err := rows.Err(); err != nil {
			log.Println(err.Error())
		}

	}
	return result, nil
}

func (c *CommentRepo) UpdateComment(obj *models.Comment)(*models.Comment, error){
	query := `
		UPDATE comments
		SET
			title = COALESCE($2, title),
			content = COALESCE($3, content),
			created_by = COALESCE($4, created_by)
		where id = $1
		RETURNING id, title, content, created_at, created_by;
	`

	var result models.Comment
	row := c.db.QueryRow(
		query,
		obj.ID,
		obj.Title,
		obj.Content,
		obj.CreatedBy,
	)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.Content,
		&result.CreatedBy,
		&result.CreatedAt,
	)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &result, nil
}

func (c *CommentRepo)GetCommentsByUserId(id int) ([]*models.Comment, error){
	query := `
		SELECT id, title, content, created_by, created_at from comments where created_by = $1
	`

	rows, err := c.db.Query(query, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	var res []*models.Comment

	for rows.Next(){
		var obj models.Comment
		if err := rows.Scan(&obj.ID, &obj.Title, &obj.Content, &obj.CreatedBy, &obj.CreatedAt); err != nil {
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


