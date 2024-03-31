package handlers

import (
	"goth/model"
	"goth/view"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	categories := []model.Category{
		{
			Name: "Model",
			Fields: []model.Field{
				{
					Name: "i5",
				},
				{
					Name: "i7",
				},
				{
					Name: "i9",
				},
			},
		},
	}
	/*
		processors := []model.Processor{
			{
				Name:         "Intel® Core™ i5-13400F, processeur pour PC de bureau, 10 cœurs (6 P-cores + 4 E-cores) 20 Mo de cache, jusqu'à 4,6 GHz",
				Price:        "199€",
				Model:        "i5",
				SocketType:   "LGA 1700",
				Architecture: "x86_64",
				Power:        "65 Watts",
				Frequency:    "2.8Ghz",
				CoreNumber:   "10",
				Link:         "https://amzn.to/3PClhzz",
			},
		}
	*/
	accessories := []model.Accessory{
		{
			Name:  "StarLinker Câble Thunderbolt 4 [2m] tressé, charge rapide, prend en charge le transfert de données 40gbps et un écran HD 8K, USB C pour Thunderbolt 3, usb4 / 3, ordinateur portable, egpu, moniteur",
			Link:  "https://amzn.to/4aAhuuI",
			Price: "22.99 €",
		},
		{
			Name:  "Mr.Kaplan Adaptateur Compact USB C Femelle vers USB Mâle 480Mb USB 2.0",
			Link:  "https://amzn.to/3TWuzZw",
			Price: "1.99 €",
		},
	}

	err := view.Layout("Components Price", categories, accessories).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
