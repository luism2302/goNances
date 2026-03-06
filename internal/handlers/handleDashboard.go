package handlers

import (
	"log"
	"net/http"

	"github.com/luism2302/goNances/views/base"
)

func (cfg *Config) HandleDashboard(w http.ResponseWriter, r *http.Request) error {
	log.Println(r.Cookies())
	return renderTemplate(w, r, base.LayoutDashboard(cfg.CurrentUser))
}
