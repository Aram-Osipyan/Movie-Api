package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Movie-Api/handlers"
	"github.com/Movie-Api/middlewares"
	_ "github.com/lib/pq"
)

func main() {
	run_init_sql()

	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/api/users", &handlers.UserHandler{})
	mux.Handle("/api/users/login", &handlers.LoginHandler{})
	mux.Handle("/api/artists", middlewares.AuthMiddleware(&handlers.ArtistHandler{}))
	mux.Handle("/api/artists/{id}", middlewares.AuthMiddleware(&handlers.ArtistHandler{}))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("jopa"))
	})
	// Run the server
	http.ListenAndServe(":8080", mux)
}

func run_init_sql() {
	connStr := os.Getenv("DB_DSN")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	// Split the data into individual statements
	statements := strings.Split(InitSqlSchemaMigration, ";")

	// Execute each statement
	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			fmt.Println("Error executing statement:", err)
		} else {
			fmt.Println("Statement executed successfully.")
		}
	}
}
