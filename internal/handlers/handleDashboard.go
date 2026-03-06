package handlers

import (
	"errors"
	"net/http"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) error {
	_, err := r.Cookie("Authorization")
	if err != nil {
		return errors.New("Couldn't find Authorizaiton cookie")
	}
	return nil
}
