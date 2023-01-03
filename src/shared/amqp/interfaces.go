package sql

import (
	"github.com/streadway/amqp"
)

type AmqpMessageBroker interface {
	Connect() (*amqp.Connection, error)
}
