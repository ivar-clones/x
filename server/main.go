package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"x/pkg/controllers"
	"x/pkg/repository"
	"x/pkg/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	repo := repository.New(conn)

	userService := user.New(repo)

	controllers := controllers.New(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", controllers.GetAllUsers)
	mux.HandleFunc("POST /api/v1/users", controllers.CreateUser)
	
	log.Println("Server started on port 3000")
	
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Printf("Unable to start server: %v\n", err)
		os.Exit(1)
	}
}