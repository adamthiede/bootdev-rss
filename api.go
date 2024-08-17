package main

import (
	"encoding/json"
	"fmt"
	"github.com/adamthiede/bootdev-rss/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	response := statusResponse{
		Status: "ok",
	}
	respondWithJSON(w, http.StatusOK, response)
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Printf("Error decoding parameters: %s\n", err)
		respondWithError(w, http.StatusBadRequest, "Error decoding parameters")
		return
	}
	//api_key := "data"
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		errtxt := fmt.Sprintf("couldn't create user '%s': \n%s\n", params.Name, err)
		fmt.Printf(errtxt)
		respondWithError(w, 400, errtxt)
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))

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
