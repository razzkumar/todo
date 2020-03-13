package queue

import (
	//"fmt"
	"os"

	"github.com/razzkumar/todo/db-service/utils/constants"
	"github.com/razzkumar/todo/db-service/utils/logger"
	"github.com/streadway/amqp"
	//"log"
)

func Connect() *amqp.Connection {

	// Get the connection string from the environment variable
	AMQP_URL := os.Getenv("AMQP_URL")

	//If it doesn't exist, use the default connection string.

	if AMQP_URL == "" {
		AMQP_URL = constants.AMQP_URL
	}
	conn, err := amqp.Dial(AMQP_URL)
	if err != nil {
		logger.FailOnError(err, "Failed to connect to RabbitMQ")
	}
	return conn
}

func CreateQueueAndBind(ch *amqp.Channel, queueName string, exchange string) {
	_, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logger.FailOnError(err, "Failed to declare a queue")
	}

	//Binding the queue to exchange data
	err = ch.QueueBind(queueName, "#", exchange, false, nil)

	if err != nil {
		logger.FailOnError(err, "Failed to Bind queue")
	}

}
