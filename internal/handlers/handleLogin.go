package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/luism2302/goNances/database/sqlc"
	"github.com/luism2302/goNances/internal/auth"
	"github.com/luism2302/goNances/views/login"
	"github.com/luism2302/goNances/views/models"
)

const (
	n = 20
)

func (cfg *Config) HandleLogin(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")

	params := models.NewLoginParams(username, password)
	errs := make(map[string]string)
	errs = params.Validate()
	if len(errs) > 0 {
		return renderTemplate(w, r, login.LoginForm(params, errs))
	}
	user, err := cfg.Queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		errs["username"] = "User not found"
		return renderTemplate(w, r, login.LoginForm(params, errs))
	}

	matches, err := auth.CheckHashedPassword(password, user.HashedPassword)
	if err != nil {
		return err
	}

	if !matches {
		errs["password"] = "Wrong Password"
		return renderTemplate(w, r, login.LoginForm(params, errs))
	}

	token, err := auth.GenerateToken(n)
	assignTokenParams := sqlc.AssignTokenToUserParams{
		SessionToken: pgtype.Text{String: token, Valid: true},
		ID:           user.ID,
	}

	if err := cfg.Queries.AssignTokenToUser(context.Background(), assignTokenParams); err != nil {
		return errors.New("Couldn't assign session token to user")
	}

	cookie := &http.Cookie{Name: "Authorization", Value: token, HttpOnly: true, Secure: true, MaxAge: 3600}
	http.SetCookie(w, cookie)
	w.Header().Set("HX-Redirect", "/dashboard")
	return nil

}
