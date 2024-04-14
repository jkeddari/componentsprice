package main

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"os"

	"github.com/jkeddari/componentsprice/handlers"
	"github.com/jkeddari/componentsprice/model"
	"github.com/joho/godotenv"
)

var data = model.Data{
	SourceList:  []string{"amazon.fr", "amazon.com"},
	ProductType: []string{"CPU", "GPU", "Accessory"},
	Items: map[string]map[string]model.Items{
		"amazon.fr": {
			"CPU": {
				Name:       "CPU",
				FieldsList: []string{"Brand", "Model", "SocketType", "Architecture", "Frequency", "Core"},
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
							"Frequency":    "2.8Ghz",
							"Core":         "10",
						},
					},
					{
						Name:  "AMD Ryzen 5 4500 avec Ventilateur Wraith Stealth - (Socket AM4/6 Cœurs -12 Threads/Frequence Min 3,6GHZ- Frequence Boost 4,1GHz/11MB/65W) - 100-100000644BOX Multicolore",
						Price: "85€",
						URL:   "https://amzn.to/49Dmhur",
						Fields: map[string]string{
							"Brand":        "AMD",
							"Model":        "Ryzen 5",
							"SocketType":   "AM4",
							"Architecture": "amd64",
							"Frequency":    "3.6 GHz",
							"Core":         "6",
						},
					},
					{
						Name:  "Intel® Core™ i5-12400F, processeur pour PC de bureau 18 Mo de cache, jusqu'à 4,40 GHz",
						Price: "142€",
						URL:   "https://amzn.to/4cXMmHc",
						Fields: map[string]string{
							"Brand":        "Intel",
							"Model":        "i5",
							"SocketType":   "LGA 1700",
							"Architecture": "x86_64",
							"Frequency":    "2.5 GHz",
							"Core":         "6",
						},
					},
					{
						Name:  "AMD Ryzen 5 5500 avec ventilateur Wraith Stealth - (socket AM4/6 Cœurs- 12 threads /frequence min 3,6GHZ - frequence boost 4,2GHz/19MB/65W) - 100-100000457BOX",
						Price: "109€",
						URL:   "https://amzn.to/3TW0059",
						Fields: map[string]string{
							"Brand":        "AMD",
							"Model":        "Ryzen 7",
							"SocketType":   "AM4",
							"Architecture": "amd64",
							"Frequency":    "3.6 GHz",
							"Core":         "6",
						},
					},
					{
						Name:  "AMD Processeur Ryzen 7 5800X Socket AM4 (3,8 Ghz) (sans iGPU) ,Noir",
						Price: "207€",
						URL:   "https://amzn.to/3UhJYEf",
						Fields: map[string]string{
							"Brand":        "AMD",
							"Model":        "Ryzen 7",
							"SocketType":   "AM4",
							"Architecture": "amd64",
							"Frequency":    "3.8 GHz",
							"Core":         "8",
						},
					},
					{
						Name:  "Intel® Core™ i5-12600KF, processeur pour PC de bureau, 10 cœurs (6P+4E) jusqu'à 4,9 GHz, LGA1700, chipset série 600 125 W",
						Price: "192.88€",
						URL:   "https://amzn.to/3VXbvf5",
						Fields: map[string]string{
							"Brand":        "Intel",
							"Model":        "i5",
							"SocketType":   "LGA 1151",
							"Architecture": "x86_64",
							"Frequency":    "3.7 GHz",
							"Core":         "6",
						},
					},
					{
						Name:  "Intel® Core™ i7-14700, processeur pour PC de bureau, 20 cœurs (8 P-cores + 12 E-cores) jusqu'à 5,4 GHz",
						Price: "430€",
						URL:   "https://amzn.to/3vT5LbQ",
						Fields: map[string]string{
							"Brand":        "Intel",
							"Model":        "i7",
							"SocketType":   "",
							"Architecture": "x86_64",
							"Frequency":    "5.4 GHz",
							"Core":         "8",
						},
					},
					{
						Name:  "Intel® Core™ i9-12900, processeur pour PC de bureau, 30 Mo de cache, jusqu'à 5,10 GHz",
						Price: "543€",
						URL:   "https://amzn.to/49DvXoF",
						Fields: map[string]string{
							"Brand":        "Intel",
							"Model":        "i9",
							"SocketType":   "LGA 1700",
							"Architecture": "x86_64",
							"Frequency":    "2.4 GHz",
							"Core":         "16",
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
					{
						Name:  "Gigabyte NVIDIA GeForce RTX 4060 GAMING OC Carte graphique - 8GB GDDR6, 128-bit, PCI-E 4.0, 2550MHz Core Clock, 2x DP 1.4, 2x HDMI 2.1a, NVIDIA DLSS 3 - GV-N4060GAMING OC-8GD",
						Price: "369.8€",
						URL:   "https://amzn.to/4cYjoaa",
						Fields: map[string]string{
							"Brand":     "GIGABYTE",
							"Memory":    "8 Go",
							"Frequency": "2550 MHz",
						},
					},
					{
						Name:  "GIGABYTE RX6600 EAGLE-8GD GV-R66EAGLE-8GD gddr6 Noir, AMD Radeon RX 6600",
						Price: "234.89€",
						URL:   "https://amzn.to/4dabcUL",
						Fields: map[string]string{
							"Brand":     "GIGABYTE",
							"Memory":    "8 Go",
							"Frequency": "2044 MHz",
						},
					},
					{
						Name:  "MSI GeForce RTX 4060 Ventus 2X Black 8G OC",
						Price: "329€",
						URL:   "https://amzn.to/440imX7",
						Fields: map[string]string{
							"Brand":     "MSI",
							"Memory":    "8 Go",
							"Frequency": "2505 MHz",
						},
					},
					{
						Name:  "ASUS NVIDIA GeForce GT 730 SL-2GD5-BRK - Carte graphique (2GB GDDR5, PCIe 2.0, Refroidissement silencieux 0 dB, GPU Tweak II avec XSplit Gamecaster, Auto-Extreme, idéal pour HTPC)",
						Price: "77.95€",
						URL:   "https://amzn.to/3xxtMFY",
						Fields: map[string]string{
							"Brand":     "ASUS",
							"Memory":    "2048 MB",
							"Frequency": "902 MHz",
						},
					},
					{
						Name:  "MSI NVIDIA GeForce RTX 3060 Ventus 2X 12G OC Noir",
						Price: "319.99€",
						URL:   "https://amzn.to/3vYlAhi",
						Fields: map[string]string{
							"Brand":     "MSI",
							"Memory":    "12 Go",
							"Frequency": "1807 MHz",
						},
					},
					{
						Name:  "MAXSUN Cartes Graphiques AMD Radeon RX 550 4Go GDDR5 Carte Graphique ITX GPU 128 Bits DirectX 12 PCI Express X16 3.0 DVI-D Dual Link, HDMI, DisplayPort",
						Price: "109.99€",
						URL:   "https://amzn.to/3TYj9n9",
						Fields: map[string]string{
							"Brand":     "maxsun",
							"Memory":    "4 Go",
							"Frequency": "1183 MHz",
						},
					},
					{
						Name:  "Gigabyte NVIDIA GeForce RTX 4060 AERO OC Carte graphique - 8GB GDDR6, 128-bit, PCI-E 4.0, 2550MHz Core Clock, 2x DP 1.4, 2x HDMI 2.1a, NVIDIA DLSS 3 - GV-N4060AERO OC-8GD",
						Price: "369.99€",
						URL:   "https://amzn.to/3Q0InzY",
						Fields: map[string]string{
							"Brand":     "GIGABYTE",
							"Memory":    "8 Go",
							"Frequency": "2550 MHz",
						},
					},
					{
						Name:  "XFX Speedster SWFT 210 Radeon RX 6600 Core Gaming Carte Graphique avec 8GB GDDR6 HDMI 3xDP, AMD RDNA™ 2 (RX-66XL8LFDQ)",
						Price: "229€",
						URL:   "https://amzn.to/3Ug6Un0",
						Fields: map[string]string{
							"Brand":     "XFX",
							"Memory":    "8 Go",
							"Frequency": "2044 MHz",
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
					{
						Name:  "Câble iPhone USB C 3M [Certifié Apple MFi], Câble USB C vers Lightning 3M Long Nylon Cable iPhone Charge Rapide Cable Lightning USB C Cordon Fil Chargeur pour iPhone 14 Pro Max/14 Plus/13/12/11/X/8/SE",
						URL:   "https://amzn.to/3xBIICY",
						Price: "9.99 €",
					},
					{
						Name:  "Station d'accueil 6 en 1 Steam Deck, AYCLIF Steam Deck Dock avec 3*USB 3.0, 1 Gbps Ethernet, HDMI 4K@60Hz, 100W PD -Chargement, Steamdeck Stand-Accessoire pour Valve Steam Deck",
						URL:   "https://amzn.to/4aDOanl",
						Price: "35.99€",
					},
					{
						Name:  "USB C Hub, AYCLIF 4-Port USB 3.2 Hub, USB C Splitter Multiport Adaptateur Data Hub 10 Gbps avec câble étendu de 50 cm pour MacBook Air/Pro, iMac, iPad Pro, Dell, HP, Surface et Appareil USB C",
						URL:   "https://amzn.to/440fSID",
						Price: "15.99€",
					},
					{
						Name:  "Station d'accueil USB C 6 in1 AYCLIF, Adattatore HDMI VGA Dual Monitor - USB 2.0, lettore schede SD/TF - Compatibile Con MacBook Pro/Air, Dell/HP/Lenovo",
						URL:   "https://amzn.to/3VTtAec",
						Price: "15.99€",
					},
					{
						Name:  "CAKOBLE Câble d'extension Gen2 1M,20Gbps 100W/5A Rallonge Mâle et Femelle Compatible USB C 3.2, USB 3.1, USB 3.0, 4K@@60Hz pour Hub/Dell XPS/USB C Charger",
						URL:   "https://amzn.to/3w2pO7A",
						Price: "14.99€",
					},
					{
						Name:  "CAKOBLE rallonge usb c 2m,usb c 3.2 gen 2 20Gbps 100W/5A Rallonge USB C Mâle et Femelle Compatible USB C 3.2/3.1/3.0, 4K@60Hz Cable C, USB C Hub/Dell XPS/USB C Charger cable usb c coudé",
						URL:   "https://amzn.to/3VYWQQW",
						Price: "15.29€",
					},
					{
						Name:  "Glangeh Support Ordinateur Portable Pliable, Support PC Portable Aluminium Ergonomique, Laptop Stand Compatible avec MacBook Air Pro, Dell XPS, HP et Plus d'Ordinateurs Portables 10-16 Pouces",
						URL:   "https://amzn.to/3vHmyPb",
						Price: "26.99€",
					},
					{
						Name:  "Vinabo Jeu Tournevis de Précision 32 en 1- Mini Set Tournevis Magnétique Bricolage Electronique de Réparation Kit, Outil Informatique pour PC, MacBook, iPhone, Lunettes, Montre, Smartphone",
						URL:   "https://amzn.to/3U0ufIl",
						Price: "8.39€",
					},
				},
			},
		},
	},
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Info("no .env file")
	}

	port := os.Getenv("SERVER_PORT")

	r := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("GET /static/*", http.StripPrefix("/static/", fileServer))

	r.HandleFunc("GET /", handlers.NewHomeHandler(&data).ServeHTTP)
	r.HandleFunc("GET /about", handlers.NewAboutHandler().ServeHTTP)
	r.HandleFunc("GET /contact", handlers.NewContactHandler().ServeHTTP)
	r.HandleFunc("GET /products", handlers.NewProductsHandler(&data).ServeHTTP)

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	certPath := os.Getenv("CERTIFICATE_PATH")
	certificateKeyPath := os.Getenv("CERTIFICATE_KEY_PATH")
	certs, err := tls.LoadX509KeyPair(certPath, certificateKeyPath)
	if err != nil {
		logger.Info("error loading certificates tls disabled", err)
	} else {
		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{certs},
		}

	}

	logger.Info("Server starting", slog.String("port", port))

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
