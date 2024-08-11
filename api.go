package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
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

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	newUser := User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Printf("Error decoding parameters: %s\n", err)
		w.WriteHeader(500)
		respondWithError(w, http.StatusBadRequest, "Error decoding parameters")
	}
	w.WriteHeader(201)

	apiCfg.DB.CreateUser()
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
