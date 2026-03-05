package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/views/login"
	"github.com/luism2302/goNances/views/models"
)

func HandleWelcomeLogin(w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, r, login.LoginDisplay(models.NewLoginParams("", ""), map[string]string{}))
}
