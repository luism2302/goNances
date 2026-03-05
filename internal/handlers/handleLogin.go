package handlers

import (
	"context"
	"net/http"

	"github.com/luism2302/goNances/internal/auth"
	"github.com/luism2302/goNances/views/login"
	"github.com/luism2302/goNances/views/models"
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
	w.Header().Set("HX-Redirect", "/dashboard")
	return nil
}
