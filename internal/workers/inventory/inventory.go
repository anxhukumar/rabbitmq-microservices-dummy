package inventory

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	Exchange            = "orders"
	InventoryQueue      = "order.created.inventory_worker"
	InventoryRoutingKey = "order.created"

	PaymentRoutingKey = "inventory.approved"
)

type OrderRequest struct {
	OrderID   int `json:"order_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

type InventoryResponse struct {
	Approved bool `json:"approved"`
}

func InventoryWorker(data OrderRequest, publishChan *amqp.Channel) {
	fmt.Println(data)
	// Simulate some inventory check work here and reply with boolean if order can be placed
	time.Sleep(5 * time.Millisecond)

	inventoryResponse := InventoryResponse{
		Approved: true,
	}

	err := rabbitmq.PublishJSON(context.Background(), publishChan, Exchange, PaymentRoutingKey, inventoryResponse)
	if err != nil {
		log.Printf("error while publishing json in inventory worker: %s\n", err)
		return
	}

}
