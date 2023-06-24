package main

import (
	"context"
	"fmt"

	"github.com/FatherCandle/go-rest-api-course/internal/comment"
	"github.com/FatherCandle/go-rest-api-course/internal/db"
)

// Run is going ot be responsible for the instantitation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect ot the database")
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed ot migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	cmtService.PostComment(context.Background(), comment.Comment{
		ID:     "6cda1fe9-e7f8-473d-8244-45d8947222fc",
		Slug:   "manual-test",
		Body:   "Hello World",
		Author: "Avner",
	})
	fmt.Println(cmtService.GetComment(context.Background(), "6cda1fe9-e7f8-473d-8244-45d8947222fc"))


	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
