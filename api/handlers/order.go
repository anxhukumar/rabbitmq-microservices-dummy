package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type OrderRequest struct {
	OrderID   int `json:"order_id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

func Order(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var orderRequest OrderRequest

	if err := decoder.Decode(&orderRequest); err != nil {
		log.Printf("error decoding json: %s\n", err)
		return
	}

	// send messages to message broker
}
