package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type CustomHandler func(http.ResponseWriter, *http.Request) error

func MakeHandler(customHandler CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := customHandler(w, r); err != nil {
			slog.Error(err.Error())
		}
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	if err := template.Render(r.Context(), w); err != nil {
		return fmt.Errorf("Couldn't render template: %w", err)
	}
	return nil
}
