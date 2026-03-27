package main

import (
	"log"
	"net/http"
	"time"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/api/handlers"
)

const port = "8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("POST /api/order", handlers.Order)

	serv := &http.Server{
		Handler:           mux,
		Addr:              ":" + port,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Server running on port: %s\n", port)
	log.Fatal(serv.ListenAndServe())
}
