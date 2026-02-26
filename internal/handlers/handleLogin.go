package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/components"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, r, components.LoginDisplay(components.LoginParams{}, map[string]string{}))
}
