package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/views/base"
)

func (cfg *Config) HandleDashboard(w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, r, base.LayoutDashboard(cfg.CurrentUser))
}
