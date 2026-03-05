package handlers

import (
	"context"
	"net/http"

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
	_, err := cfg.Queries.GetUserByUsername(context.Background(), username)
	if err != nil {
		errs["username"] = "User not found"
		renderTemplate(w, r, login.LoginForm(params, errs))
	}
	return nil
}
