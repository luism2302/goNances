package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/luism2302/goNances/components"
	"github.com/luism2302/goNances/database/sqlc"
	"github.com/luism2302/goNances/internal/auth"
)

func (cfg *Config) HandleUsersCreate(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confPassword")

	_, err := cfg.Queries.GetUserByUsername(context.Background(), username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	userExists := err == nil
	params := components.NewSignUpParams(username, password, confirmPassword, userExists)

	errors := params.Validate()

	if len(errors) > 0 {
		err := renderTemplate(w, r, components.SignUpForm(params, errors))
		return err
	}

	hashed, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	newUserParams := sqlc.CreateUserParams{Username: username, HashedPassword: hashed}
	_, err = cfg.Queries.CreateUser(context.Background(), newUserParams)
	if err != nil {
		return fmt.Errorf("Couldn't create new user: %w", err)
	}
	w.Header().Set("HX-Redirect", "/")
	return nil
}
