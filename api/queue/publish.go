package queue

import (
	"fmt"

	"github.com/razzkumar/todo/api/utils/constants"
	"github.com/razzkumar/todo/api/utils/logger"
	"github.com/streadway/amqp"
	//"log"
)

func Connect() *amqp.Connection {
	conn, err := amqp.Dial(constants.AMQP_URL)

	if err != nil {
		logger.FailOnError(err, "Failed to connect to RabbitMQ")
	}
	fmt.Println("----------AMQP-Connected------------------")
	return conn
}

func CreateQueue(ch *amqp.Channel, queueName string) {
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
	//err = ch.QueueBind(queueName, "#", exchange, false, nil)

	//if err != nil {
	//logger.FailOnError(err, "Failed to Bind queue")
	//}

}
