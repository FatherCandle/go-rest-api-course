package main

import (
	"context"
	"fmt"

	"github.com/FatherCandle/go-rest-api-course/internal/db"
)

// Run is going ot be responsible for the instantitation and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect ot the database")
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil{
		fmt.Println(err)
	}
}