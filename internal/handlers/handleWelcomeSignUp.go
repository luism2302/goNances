package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/views/models"
	"github.com/luism2302/goNances/views/signup"
)

func HandleWelcomeSignUp(w http.ResponseWriter, r *http.Request) error {
	if err := renderTemplate(w, r, signup.SignUpDisplay(models.SignUpParams{}, map[string]string{})); err != nil {
		return err
	}
	return nil
}
