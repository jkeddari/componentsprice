package handlers

import (
	"goth/model"
	"goth/view"
	"log"
	"net/http"
)

type HomeHandler struct {
	data *model.Data
}

func NewHomeHandler(d *model.Data) *HomeHandler {
	return &HomeHandler{d}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := view.Layout("Components Price", *h.data).Render(r.Context(), w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
