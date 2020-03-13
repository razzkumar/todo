package main

import (
	//"fmt"

	"github.com/gin-gonic/gin"
	"github.com/razzkumar/todo/api/handler"
	"github.com/razzkumar/todo/api/models"
	"github.com/razzkumar/todo/api/queue"
	"github.com/razzkumar/todo/api/utils/constants"
	"github.com/razzkumar/todo/api/utils/logger"
	//"github.com/streadway/amqp"
)

func main() {
	//fmt.Println("Hello world")

	// Connect rabbit mq and Create
	conn := queue.Connect()

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		logger.FailOnError(err, "Fail to create channel")
	}

	defer ch.Close()

	// We create an exahange that will bind to the queue to send and receive messages
	//err = ch.ExchangeDeclare(constants.EXCHANGER, "topic", true, false, false, false, nil)

	//if err != nil {
	//logger.FailOnError(err, "Fail to Declare Exchage")
	//}

	//queue.CreateQueueAndBind(ch, constants.QUEUE_NAME, constants.EXCHANGER)
	queue.CreateQueue(ch, constants.QUEUE_NAME)

	todo := models.New()
	todo.Add(
		models.Todo{
			ID:     1,
			Data:   "this is default todo 1",
			IsDone: true,
		},
	)
	todo.Add(
		models.Todo{
			ID:     2,
			Data:   "this is default todo 2",
			IsDone: false,
		},
	)

	router := gin.Default()

	router.GET("/", handler.Hello)
	router.GET("/api/todos", handler.GetTodos(todo))
	router.POST("/api/todos", handler.AddTodo(todo, ch))
	router.Run(":8888")
}
