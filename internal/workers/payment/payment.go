package payment

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	Exchange          = "orders"
	PaymentQueue      = "inventory.approved.payment_worker"
	PaymentRoutingKey = "inventory.approved"

	NotificationRoutingKey = "payment.confirmed"
)

type InventoryApprovalRequest struct {
	Approved bool `json:"approved"`
}

type PaymentResponse struct {
	Status string `json:"status"`
}

func PaymentWorker(data InventoryApprovalRequest, publishChan *amqp.Channel) {
	fmt.Println(data)
	// Simulate payment processing
	time.Sleep(50 * time.Millisecond)

	paymentResponse := PaymentResponse{
		Status: "payment processed",
	}

	err := rabbitmq.PublishJSON(context.Background(), publishChan, Exchange, NotificationRoutingKey, paymentResponse)
	if err != nil {
		log.Printf("error while publishing json in inventory worker: %s\n", err)
		return
	}
}
