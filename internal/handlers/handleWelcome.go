package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/components"
)

func HandleWelcome(w http.ResponseWriter, r *http.Request) error {
	if err := renderTemplate(w, r, components.LayoutWelcome()); err != nil {
		return err
	}
	return nil
}
