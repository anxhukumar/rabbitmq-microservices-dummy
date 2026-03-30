package main

import (
	"log"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	rabbitmqConnURL        = "amqp://guest:guest@localhost:5672/"
	orderExchange          = "orders.exchange"
	orderCreatedQueue      = "inventory.order.queue"
	orderCreatedRoutingKey = "order.created.key"
)

func main() {
	// amqp: connection
	amqpConn, err := amqp.Dial(rabbitmqConnURL)
	if err != nil {
		log.Println("rabbitmq connection failed in inventory")
		return
	}
	defer amqpConn.Close()
	log.Println("connection to rabbitmq successful in inventory")

	// Declare and bind queue
	_, _, err = rabbitmq.DeclareAndBind(
		amqpConn,
		orderExchange,
		orderCreatedQueue,
		orderCreatedRoutingKey,
		rabbitmq.Transient,
	)
	if err != nil {
		log.Printf("error while declaring and binding queue in inventory: %s\n", err)
		return
	}
}
