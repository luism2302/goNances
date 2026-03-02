package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/components"
)

func HandleUsersCreate(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confPassword")

	params := components.NewSignUpParams(username, password, confirmPassword)
	errors := params.Validate()

	if len(errors) > 0 {
		err := renderTemplate(w, r, components.SignUpForm(params, errors))
		return err
	}
	//TODO: create new user in db
	w.Header().Set("HX-Redirect", "/")
	return nil
}
