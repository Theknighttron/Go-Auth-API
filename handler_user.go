package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/polyhistor2050/Go-Auth-API/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Fistname  string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Gender    string `json:"gender"`
		StudentId string `json:"studentId"`
		Course    string `json:"course"`
		Level     string `json:"level"`
		Email     string `json:"email"`
	}

	// parse the request body into the struct
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Create CreateUserParams struct and populate it
	createUserParams := database.CreateUserParams{
		Firstname: params.Fistname,
		Lastname:  params.Lastname,
		Gender:    params.Gender,
		StudentId: params.StudentId,
		Course:    params.Course,
		Level:     params.Level,
		Email:     params.Email,
	}

	// Call CreateUser function with the provided parameters
	user, err := apiCfg.DB.CreateUser(r.Context(), createUserParams)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 200, user)
}
