package main

import (
	"goth/handlers"
	"log/slog"
	"net/http"
	"os"
)

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

	r.HandleFunc("GET /", handlers.NewHomeHandler().ServeHTTP)
	r.HandleFunc("GET /about", handlers.NewAboutHandler().ServeHTTP)

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
