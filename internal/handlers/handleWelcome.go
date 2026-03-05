package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/views/base"
)

func HandleWelcome(w http.ResponseWriter, r *http.Request) error {
	if err := renderTemplate(w, r, base.LayoutWelcome()); err != nil {
		return err
	}
	return nil
}
