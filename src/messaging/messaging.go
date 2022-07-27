package messaging

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Ctx = context.TODO()

func connectQueueChannel(onConnect func(*amqp.Channel, *amqp.Queue) error) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
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
		"petcare-work", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return err
	}
	return onConnect(ch, &q)
}

func Emit(body string) error {
	return connectQueueChannel(func(ch *amqp.Channel, q *amqp.Queue) error {
		return ch.PublishWithContext(
			Ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
	})
}

func Consume(consume func(string)) error {
	return connectQueueChannel(func(ch *amqp.Channel, q *amqp.Queue) error {
		err := ch.Qos(
			1,     // prefetch count
			0,     // prefetch size
			false, // global
		)
		if err != nil {
			return err
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			return err
		}

		go func() {
			for msg := range msgs {
				consume(string(msg.Body))
			}
		}()

		select {}
	})
}
