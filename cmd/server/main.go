package main

import (
	"fmt"

	"github.com/FatherCandle/go-rest-api-course/internal/comment"
	transportHttp "github.com/FatherCandle/go-rest-api-course/internal/transport/http"
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

	httpsHandler := transportHttp.NewHandler(cmtService)
	if err := httpsHandler.Serve(); err != nil {
		return err
	}


	return nil
}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
