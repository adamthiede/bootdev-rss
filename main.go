package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}
	listenOn := fmt.Sprintf("%s:%s", addr, port)
	fmt.Printf("Listening on %s\n", listenOn)

	sm := http.NewServeMux()
	server := http.Server{
		Handler: sm,
		Addr:    listenOn,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
