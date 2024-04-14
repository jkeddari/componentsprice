package handlers

import (
	"log"
	"net/http"

	"github.com/jkeddari/componentsprice/model"
	"github.com/jkeddari/componentsprice/views/page"
)

type HomeHandler struct {
	data *model.Data
}

func NewHomeHandler(d *model.Data) *HomeHandler {
	return &HomeHandler{d}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := page.Home("Components Price", *h.data).Render(r.Context(), w)

	log.Println("new connection :", r.RemoteAddr, " | on agent :", r.UserAgent())

	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
