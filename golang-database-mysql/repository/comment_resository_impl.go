package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context,comment entity.Comment) (entity.Comment, error) {
	query := "Insert into comments (email, comment) values (?, ?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
		
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context,id int32) (entity.Comment, error) {
	query := "Select id, email, comment from comments where id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
	} else {
		return comment, errors.New("Comment Not Found")
	}
	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "Select id, email, comment from comments"
	rows, err := repository.DB.QueryContext(ctx, query)
	comments := []entity.Comment{}
	if err != nil {
		return comments, err
	}
	defer rows.Close()
	for rows.Next() {
		comment := entity.Comment{}
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}