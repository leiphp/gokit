package mq

import (
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func Init(uri string) error {
	var err error
	conn, err = amqp.Dial(uri)
	if err != nil {
		return err
	}
	ch, err = conn.Channel()
	return err
}

func Publish(queue string, body []byte) error {
	_, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	return ch.Publish("", queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
}

func Consume(queue string, handler func([]byte)) error {
	_, err := ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}
	deliveries, err := ch.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range deliveries {
			handler(d.Body)
		}
	}()
	return nil
}
