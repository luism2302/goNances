package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/luism2302/goNances/database/sqlc"
	"github.com/luism2302/goNances/internal/auth"
	"github.com/luism2302/goNances/views/models"
	"github.com/luism2302/goNances/views/signup"
)

func (cfg *Config) HandleUsersCreate(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confPassword")
	errs := make(map[string]string)

	params := models.NewSignUpParams(username, password, confirmPassword)
	errs = params.Validate()
	if len(errs) > 0 {
		err := renderTemplate(w, r, signup.SignUpForm(params, errs))
		return err
	}

	_, err := cfg.Queries.GetUserByUsername(context.Background(), username)
	if err == nil {
		errs["username"] = "User already exists"
		return renderTemplate(w, r, signup.SignUpForm(params, errs))
	}

	if !errors.Is(err, sql.ErrNoRows) {
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
