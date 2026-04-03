package notification

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	Exchange               = "orders"
	NotificationQueue      = "payment.confirmed.notification_worker"
	NotificationRoutingKey = "payment.confirmed"
)

type PaymentStatusRequest struct {
	Status string `json:"status"`
}

func NotificationWoker(data PaymentStatusRequest, publishChan *amqp.Channel) {
	fmt.Println(data)
	// Simulate notification being sent to user
	time.Sleep(10 * time.Millisecond)

	log.Printf("Payment Cycle completed...")
}
