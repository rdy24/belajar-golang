package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestRepository(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "tess123@mail.com",
		Comment: "tes123",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
func TestRepositoryFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
func TestRepositoryFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _,comment := range comments {
		fmt.Println(comment)
	}
}