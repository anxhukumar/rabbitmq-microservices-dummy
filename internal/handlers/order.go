package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anxhukumar/rabbitmq-microservices-dummy/internal/rabbitmq"
)

type OrderRequest struct {
	OrderID   int `json:"order_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

const (
	exchnage   = "orders"
	routingKey = "order.created"
)

func (s *Server) Order(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var orderRequest OrderRequest

	if err := decoder.Decode(&orderRequest); err != nil {
		log.Printf("error decoding json: %s\n", err)
		return
	}

	// send messages to message broker
	err := rabbitmq.PublishJSON(r.Context(), s.PublishCh, exchnage, routingKey, orderRequest)
	if err != nil {
		log.Printf("error publishing json to rabbitmq: %s", err)
		return
	}

}
