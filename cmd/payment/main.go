package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const rabbitmqConnURL = "amqp://guest:guest@localhost:5672/"

func main() {
	// amqp: connection
	amqpConn, err := amqp.Dial(rabbitmqConnURL)
	if err != nil {
		log.Println("rabbitmq connection failed in payment")
		return
	}
	defer amqpConn.Close()
	log.Println("connection to rabbitmq successful in payment")
}
