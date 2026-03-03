package handlers

import (
	"net/http"
	"context"
	"fmt"
	"log"

	"github.com/luism2302/goNances/components"
	"github.com/luism2302/goNances/database/sqlc"
)

func (cfg *Config) HandleUsersCreate(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confPassword")

	params := components.NewSignUpParams(username, password, confirmPassword)
	errors := params.Validate()

	if len(errors) > 0 {
		err := renderTemplate(w, r, components.SignUpForm(params, errors))
		return err
	}
	//TODO: create new user in db
	newUserParams := sqlc.CreateUserParams{Username: username, Password: password}
	user, err := cfg.Queries.CreateUser(context.Background(), newUserParams)
	if err != nil {
		fmt.Errorf("Couldn't create new user: %w", err)
	}
	log.Print(user)
	w.Header().Set("HX-Redirect", "/")
	return nil
}
