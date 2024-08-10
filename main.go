package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type statusResponse struct {
	Status string `json:"status"`
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "127.0.0.1"
	}
	listenOn := fmt.Sprintf("%s:%s", addr, port)
	fmt.Printf("Listening on %s\n", listenOn)

	smux := http.NewServeMux()

	// healthz
	healthzHandler := func(w http.ResponseWriter, r *http.Request) {
		response := statusResponse{
			Status: "ok",
		}
		respondWithJSON(w, http.StatusOK, response)
	}
	smux.HandleFunc("GET /v1/healthz", healthzHandler)

	// error
	errHandler := func(w http.ResponseWriter, r *http.Request) {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	smux.HandleFunc("GET /v1/err", errHandler)

	// run http server after every handler is added
	server := http.Server{
		Handler: smux,
		Addr:    listenOn,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
	fmt.Printf("Error marshalling JSON: %s\n", err)
	w.WriteHeader(500)
	return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	fmt.Printf("responding with %v: %s\n", code, msg)
	type errResp struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResp{
		Error: msg,
	})
}
