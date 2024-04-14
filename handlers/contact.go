package handlers

import (
	"net/http"

	"github.com/jkeddari/componentsprice/view/page"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

func (h *ContactHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := page.Contact("Contact - Componentsprice").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
