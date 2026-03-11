package handlers

import (
	"net/http"

	"github.com/luism2302/goNances/views/components"
	"github.com/luism2302/goNances/views/models"
)

func HandleExpenses(w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, r, components.ExpensesContent(models.NewExpenseParams(0, "", ""), map[string]string{}))
}
