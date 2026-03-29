package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
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
