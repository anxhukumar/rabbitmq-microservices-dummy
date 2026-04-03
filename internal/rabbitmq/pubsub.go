package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Queue types
type QueueType string

const (
	Durable   QueueType = "durable"
	Transient QueueType = "transient"
)

func PublishJSON[T any](ctx context.Context, ch *amqp.Channel, exchange, key string, val T) error {

	jsonBytes, err := json.Marshal(val)
	if err != nil {
		err := fmt.Errorf("error marshalling json: %w", err)
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		exchange,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBytes,
		},
	)
	if err != nil {
		err := fmt.Errorf("error publishing json: %w", err)
		return err
	}

	return nil
}

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType QueueType,
) (*amqp.Channel, amqp.Queue, error) {

	// Channel
	channel, err := conn.Channel()
	if err != nil {
		err := fmt.Errorf("error in creating creating connection while declaring and binding queue: %w", err)
		return nil, amqp.Queue{}, err
	}

	// Declare queue
	queue, err := channel.QueueDeclare(
		queueName,
		queueType == Durable,
		queueType == Transient,
		queueType == Transient,
		false,
		nil,
	)
	if err != nil {
		err := fmt.Errorf("error while declaring queue: %w", err)
		return nil, amqp.Queue{}, err
	}

	// Bind queue
	err = channel.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		err := fmt.Errorf("error while binding queue: %w", err)
		return nil, amqp.Queue{}, err
	}

	return channel, queue, nil
}

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType QueueType,
	handler func(T, *amqp.Channel),
) error {

	// Declare and bind queue
	channel, queue, err := DeclareAndBind(
		conn,
		exchange,
		queueName,
		key,
		queueType,
	)
	if err != nil {
		err = fmt.Errorf("error while declaring and binding queue: %w", err)
		return err
	}

	// Consume messages from queue
	delivery, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		err = fmt.Errorf("error while consuming messages: %w", err)
		return err
	}

	// Range over delivery for messages
	go func() {
		for msg := range delivery {
			var data T
			if err := json.Unmarshal(msg.Body, &data); err != nil {
				msg.Nack(false, false)
				log.Println("error:", err)
				continue
			}

			handler(data, channel)

			msg.Ack(false)
		}
	}()

	return nil
}
