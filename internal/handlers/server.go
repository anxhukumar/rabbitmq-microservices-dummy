package handlers

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Server struct {
	AmqpConn *amqp.Connection
}

func NewServer(conn *amqp.Connection) *Server {
	return &Server{
		AmqpConn: conn,
	}
}
