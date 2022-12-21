package repository

import (
	"belajar-go-mysql/entity"
	"context"
	"database/sql"
	"errors"
)

type commentRepositoryImplementation struct {
	DB *sql.DB
}

func (repository *commentRepositoryImplementation) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	//TODO implement me
	querry := "insert into comment(email, comment) values (?,?)"
	result, err := repository.DB.ExecContext(ctx, querry, comment.Email, comment.Comment)
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

func (repository *commentRepositoryImplementation) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	//TODO implement me
	querry := "select id, email, comment from comment where id = ? limit 1"
	rows, err := repository.DB.QueryContext(ctx, querry, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada data
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//jika tidak ada data
		return comment, errors.New("tidak ada data")
	}
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImplementation{db}
}

func (repository *commentRepositoryImplementation) FindAll(ctx context.Context) ([]entity.Comment, error) {
	//TODO implement me
	querry := "select id, email, comment from comment"
	rows, err := repository.DB.QueryContext(ctx, querry)
	//comment := entity.Comment{}
	if err != nil {
		return nil, err
	}
	var Comments []entity.Comment
	defer rows.Close()
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		Comments = append(Comments, comment)
	}
	return Comments, nil
}
