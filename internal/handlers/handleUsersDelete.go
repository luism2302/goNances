package handlers

import (
	"context"
	"errors"
	"net/http"
)

func (cfg *Config) HandleUsersDelete(w http.ResponseWriter, r *http.Request) error {
	if err := cfg.Queries.DeleteAllUsers(context.Background()); err != nil {
		return errors.New("Couldn't delete users")
	}
	response := DeleteResponse{Msg: "Deleted users successfully"}
	return respondJSON(w, http.StatusAccepted, response)
}
