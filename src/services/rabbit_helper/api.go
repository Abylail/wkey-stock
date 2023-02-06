package rabbit_helper

import (
	"github.com/lowl11/lazylog/layers"
	"github.com/streadway/amqp"
	"wkey-stock/src/definition"
)

func Connection(connectionURL string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(connectionURL)
	if err != nil {
		if err = connection.Close(); err != nil {
			definition.Logger.Error(err, "Close connection error", layers.Rabbit)
		}

		return nil, err
	}

	return connection, nil
}
