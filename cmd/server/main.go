package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/handlers"
	amqp "github.com/rabbitmq/amqp091-go"
)

const port = "8080"
const rabbitmqConnURL = "amqp://guest:guest@localhost:5672/"

func main() {
	// amqp: connection
	amqpConn, err := amqp.Dial(rabbitmqConnURL)
	if err != nil {
		log.Println("rabbitmq connection failed in server")
		return
	}
	defer amqpConn.Close()
	log.Println("connection to rabbitmq successful in server")

	// Create new Server
	server := handlers.NewServer(amqpConn)

	// http server
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/healthz", server.HandlerReadiness)
	mux.HandleFunc("POST /api/order", server.Order)

	serv := &http.Server{
		Handler:           mux,
		Addr:              ":" + port,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Run server in a goroutine
	go func() {
		log.Printf("server running on port: %s\n", port)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Create channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify on interrupt signals
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block here until signal is received
	sig := <-sigChan
	log.Printf("received signal: %v, shutting down...\n", sig)

	// Graceful shutdown of HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %v\n", err)
	}

	log.Println("server exited properly")
}
