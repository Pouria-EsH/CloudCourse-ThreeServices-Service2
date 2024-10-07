package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type CloudAMQ struct {
	URL       string
	QueueName string
}

func NewCloudAMQ(url string, queue string) *CloudAMQ {
	return &CloudAMQ{
		URL:       url,
		QueueName: queue,
	}
}

func (c CloudAMQ) Listen(handler func(msg string) error) error {
	conn, err := amqp.Dial(c.URL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		c.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("Consumer started on %s from %s", q.Name, c.URL)

	for d := range msgs {
		err := handler(string(d.Body))
		if err != nil {
			return fmt.Errorf("error at amq message handler: %w", err)
		}
	}

	return nil
}
