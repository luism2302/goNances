package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/components"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	if err := renderTemplate(w, r, components.Layout(components.Login(), "GoNances")); err != nil {
		return err
	}
	return nil
}
