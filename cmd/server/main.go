package main

import (
	"goth/handlers"
	"goth/model"
	"log/slog"
	"net/http"
	"os"
)

var data = model.Data{
	SourceList:  []string{"amazon.fr", "amazon.com"},
	ProductType: []string{"CPU", "GPU", "Accessory"},
	Items: map[string]map[string]model.Items{
		"amazon.fr": {
			"CPU": {
				Name:       "CPU",
				FieldsList: []string{"Brand", "Model", "SocketType", "Architecture", "Power", "Frequency", "Core"},
				Items: []model.Item{
					{
						Name:  "Intel® Core™ i5-13400F, processeur pour PC de bureau, 10 cœurs (6 P-cores + 4 E-cores) 20 Mo de cache, jusqu'à 4,6 GHz",
						Price: "205€",
						URL:   "https://amzn.to/3PClhzz",
						Fields: map[string]string{
							"Brand":        "Intel",
							"Model":        "i5",
							"SocketType":   "LGA 1700",
							"Architecture": "x86_64",
							"Power":        "65 Watts",
							"Frequency":    "2.8Ghz",
							"Core":         "10",
						},
					},
				},
			},
			"GPU": {
				Name:       "GPU",
				FieldsList: []string{"Brand", "Memory", "Frequency"},
				Items: []model.Item{
					{
						Name:  "MSI GeForce RTX 4060 Ti Ventus 2X Black 8G OC Carte Graphique - 8 Go GDDR6 (18 GB/s, 128-bit), PCIe 4.0-2 x Ventilateurs TORX Fan 4.0 - HDMI 2.1a, DisplayPort 1.4a",
						Price: "389.90€",
						URL:   "https://amzn.to/4cMK0uX",
						Fields: map[string]string{
							"Brand":     "NVIDIA",
							"Memory":    "8 Go",
							"Frequency": "2580 MHz",
						},
					},
				},
			},
			"Accessory": {
				Name:       "Accessory",
				FieldsList: []string{},
				Items: []model.Item{
					{
						Name:  "StarLinker Câble Thunderbolt 4 [2m] tressé, charge rapide, prend en charge le transfert de données 40gbps et un écran HD 8K, USB C pour Thunderbolt 3, usb4 / 3, ordinateur portable, egpu, moniteur",
						URL:   "https://amzn.to/4aAhuuI",
						Price: "22.99 €",
					},
					{
						Name:  "Mr.Kaplan Adaptateur Compact USB C Femelle vers USB Mâle 480Mb USB 2.0",
						URL:   "https://amzn.to/3TWuzZw",
						Price: "1.99 €",
					},
				},
			},
		},
	},
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	tlsServer := true
	port := os.Getenv("SERVER_PORT")

	certPath := os.Getenv("CERTIFICATE_PATH")
	certificateKeyPath := os.Getenv("CERTIFICATE_KEY_PATH")
	if certificateKeyPath == "" || certPath == "" {
		tlsServer = false
		logger.Info("tls disabled")
	}

	r := http.NewServeMux()

	r.HandleFunc("GET /", handlers.NewHomeHandler(&data).ServeHTTP)
	r.HandleFunc("GET /about", handlers.NewAboutHandler().ServeHTTP)
	r.HandleFunc("GET /products", handlers.NewProductsHandler(&data).ServeHTTP)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	logger.Info("Server starting", slog.String("port", port))
	var err error
	if tlsServer {
		err = srv.ListenAndServeTLS(certPath, certificateKeyPath)
	} else {
		err = srv.ListenAndServe()
	}

	if err != nil {
		logger.Error(err.Error())
	}
}
