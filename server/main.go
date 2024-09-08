package main

import (
	"context"
	"fmt"
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
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()


	repo := repository.New(conn)

	userService := user.New(repo)

	controllers := controllers.New(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", controllers.GetAllUsers)
}