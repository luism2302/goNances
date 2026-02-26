package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/components"
)

func HandleSignUp(w http.ResponseWriter, r *http.Request) error {
	if err := renderTemplate(w, r, components.SignUpDisplay(components.SignUpParams{}, map[string]string{})); err != nil {
		return err
	}
	return nil
}
