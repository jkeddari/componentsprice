package handlers

import (
	"goth/model"
	"goth/view"
	"net/http"
)

type ProductsHandler struct {
	Data *model.Data
}

func NewProductsHandler(data *model.Data) *ProductsHandler {
	return &ProductsHandler{Data: data}
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	productType := r.FormValue("productType")

	products := h.Data.Items[source][productType]

	err := view.ContentView(products).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
