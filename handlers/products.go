package handlers

import (
	"net/http"

	"github.com/jkeddari/componentsprice/model"
	"github.com/jkeddari/componentsprice/views/products"
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

	productsList := h.Data.Items[source][productType]

	err := products.Show(productsList).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
