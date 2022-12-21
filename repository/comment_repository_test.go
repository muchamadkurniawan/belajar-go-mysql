package repository

import (
	belajar_go_mysql "belajar-go-mysql"
	"belajar-go-mysql/entity"
	"context"
	"fmt"
	"testing"
)

import _ "github.com/go-sql-driver/mysql"

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_mysql.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@gmail.com",
		Comment: "new comment in repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("result :", result)
}

func TestFindByID(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_mysql.GetConnection())
	result, err := commentRepository.FindById(context.Background(), 36)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_go_mysql.GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
