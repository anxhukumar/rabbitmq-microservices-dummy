package handlers

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Server struct {
	PublishCh *amqp.Channel
}

func NewServer(pubChan *amqp.Channel) *Server {
	return &Server{
		PublishCh: pubChan,
	}
}
