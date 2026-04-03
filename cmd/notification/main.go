package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/rabbitmq"
	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/workers/notification"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitmqConnURL = "amqp://guest:guest@localhost:5672/"
)

func main() {
	// amqp: connection
	amqpConn, err := amqp.Dial(rabbitmqConnURL)
	if err != nil {
		log.Println("rabbitmq connection failed in notification")
		return
	}
	defer amqpConn.Close()
	log.Println("connection to rabbitmq successful in notification")

	// Create, bind and subscribe queue
	if err = rabbitmq.SubscribeJSON(
		amqpConn,
		notification.Exchange,
		notification.NotificationQueue,
		notification.NotificationRoutingKey,
		rabbitmq.Transient,
		notification.NotificationWoker,
	); err != nil {
		log.Printf("error while subscribing json in notification: %s\n", err)
		return
	}

	// Create channel to listen for OS signals
	sigChan := make(chan os.Signal, 1)

	// Notify on interrupt signals
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block here until signal is received
	<-sigChan

	log.Printf("Notification worker closed...")

}
