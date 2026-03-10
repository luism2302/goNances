package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/luism2302/goNances/database/sqlc"
)

type CustomHandler func(http.ResponseWriter, *http.Request) error

func MakeHandler(customHandler CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := customHandler(w, r); err != nil {
			slog.Error(err.Error())
		}
	}
}

func (cfg *Config) MiddlewareLoggedIn(next CustomHandler) CustomHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		authCookie, err := r.Cookie("Authorization")
		if err != nil {
			return errors.New("Couldn't find Authorizaiton cookie")
		}
		loggedUser, err := cfg.Queries.GetUserByUsername(context.Background(), cfg.CurrentUser)
		if err != nil {
			return fmt.Errorf("Couldn't find user: %s in db", cfg.CurrentUser)
		}
		if loggedUser.SessionToken.String != authCookie.Value {
			return errors.New("Couldn't authenticate user")
		}
		return next(w, r)
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	if err := template.Render(r.Context(), w); err != nil {
		return fmt.Errorf("Couldn't render template: %w", err)
	}
	return nil
}

type Config struct {
	Queries     *sqlc.Queries
	CurrentUser string
}

func NewConfig(queries *sqlc.Queries) *Config {
	return &Config{Queries: queries}
}

func respondJSON(w http.ResponseWriter, code int, payload any) error {
	marshaled, err := json.Marshal(payload)
	if err != nil {
		return errors.New("Couldn't marshal to json")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(marshaled)
	return nil
}

type DeleteResponse struct {
	Msg string `json:"msg"`
}
