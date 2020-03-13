package main

import (
	"context"
	"encoding/json"
	"fmt"
	//"time"

	"github.com/razzkumar/todo/db-service/models"
	"github.com/razzkumar/todo/db-service/mongodb"
	"github.com/razzkumar/todo/db-service/queue"
	"github.com/razzkumar/todo/db-service/utils/constants"
	"github.com/razzkumar/todo/db-service/utils/logger"
	//"go.mongodb.org/mongo-driver/bson"
)

func main() {

	/*
	 Queue
	*/

	conn := queue.Connect()
	defer conn.Close()
	ch, err := conn.Channel()
	defer ch.Close()

	/*
		MongoDB
	*/

	dbClient := mongodb.Connect()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = dbClient.Connect(ctx)

	if err != nil {
		logger.FailOnError(err, "Fail to connect mongodb")
	}

	defer dbClient.Disconnect(ctx)

	todoDb := dbClient.Database("tododb")
	todoCollection := todoDb.Collection("todos")

	// Check the connection
	err = dbClient.Ping(ctx, nil)

	if err != nil {
		logger.FailOnError(err, "DB connection error")
	}

	// We consume data in the queue named test using the channel we created in go.
	msgs, err := ch.Consume(constants.QUEUE_NAME, "", false, false, false, false, nil)

	if err != nil {
		logger.FailOnError(err, "error consuming the queue: ")
	}

	for msg := range msgs {
		//print the message to the console

		fmt.Println("message received: " + string(msg.Body))
		// Acknowledge that we have received the message so it can be removed from the queue
		todo := models.Todo{}

		//err := bson.Unmarshal(msg.Body, &todo)
		err := json.Unmarshal(msg.Body, &todo)
		if err != nil {
			logger.FailOnError(err, "Fail to Decode message")
		}

		fmt.Println(todo.ID)
		_, err = todoCollection.InsertOne(ctx, todo)
		if err != nil {
			fmt.Println("Fill to add todo")
			fmt.Println(err)
		}
		fmt.Println("Saved ")
		msg.Ack(false)
	}
}
