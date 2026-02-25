package handlers

import (
	"log"
	"net/http"
)

func HandleUsersCreate(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("conf_password")

	log.Printf("New user: %s with pswd: %s and conf: %s", username, password, confirmPassword)
	return nil
}
