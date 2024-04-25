package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Gender    string `json:"gender"`
	StudentId string `json:"studentid"`
	Course    string `json:"course"`
	Level     string `json:"level"`
	Email     string `json:"email"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate user input
	if user.Firstname == "" || user.Lastname == "" || user.Gender == "" || user.StudentId == "" || user.Course == "" || user.Level == "" || user.Email == "" {
		http.Error(w, "All fields are required to be filled", http.StatusBadRequest)
		return
	}

	// Save user to the database (assuming you have a function dbCreateUser)
	// Replace this with your actual database interaction code
	err = dbCreateUser(user.Firstname, user.Lastname, user.Gender, user.StudentId, user.Course, user.Level, user.Email)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Return success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// Dummy function to simulate database interaction
func dbCreateUser(firstname, lastname, gender, studentId, course, level, email string) error {
	// Here you would perform database operations to insert the user data
	// Replace this with actual database interaction code
	return nil
}
