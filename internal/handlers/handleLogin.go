package handlers

import (
	"log"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	log.Print("Login request made")
	return nil
}
