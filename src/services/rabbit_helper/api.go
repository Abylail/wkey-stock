package rabbit_helper

import (
	"github.com/lowl11/lazylog/log"
	"github.com/streadway/amqp"
)

func Connection(connectionURL string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(connectionURL)
	if err != nil {
		if err = connection.Close(); err != nil {
			log.Error(err, "Close RMQ connection error")
		}

		return nil, err
	}

	return connection, nil
}
