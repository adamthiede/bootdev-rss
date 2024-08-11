package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/adamthiede/bootdev-rss/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

type statusResponse struct {
	Status string `json:"status"`
}

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// load env
	godotenv.Load()
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDR")
	dbconn := os.Getenv("DBCONN")
	// set up server listening
	if addr == "" {
		addr = "127.0.0.1"
	}
	listenOn := fmt.Sprintf("%s:%s", addr, port)
	fmt.Printf("Listening on %s\n", listenOn)
	fmt.Printf("Connecting to %s\n", dbconn)
	// connect to database
	db, dberr := sql.Open("postgres", dbconn)
	if dberr != nil {
		fmt.Printf("cannot connect to database:\n%s\n%s\n", dbconn, dberr)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	smux := http.NewServeMux()

	// healthz
	smux.HandleFunc("GET /v1/healthz", healthzHandler)
	// error
	smux.HandleFunc("GET /v1/err", errHandler)
	// create user
	smux.HandleFunc("POST /v1/users", createUserHandler)

	//run http server after every handler is added
	server := http.Server{
		Handler: smux,
		Addr:    listenOn,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
