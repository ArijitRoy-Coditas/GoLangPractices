package main

import (
	"encoding/json"
	"fmt"
	"http-requests/internal/auth"
	"http-requests/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameter struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)
	if err != nil{
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v",err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v",err))
	}
	responseWithJSON(w, 201, databaseUserToUser(user))
}


func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v",err)) 
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v",err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}