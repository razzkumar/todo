package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/razzkumar/todo/api/models"
	"github.com/razzkumar/todo/api/utils/constants"
	"github.com/razzkumar/todo/api/utils/logger"
	"github.com/streadway/amqp"
)

//func AddTodo(c *gin.Context) {
//body := c.Request.Body
//val, err := ioutil.ReadAll(body)
//if err != nil {
//fmt.Println(err.Error())
//}
//c.JSON(http.StatusOK, gin.H{
//"message": string(val),
//})
//}

/*
  Helps to get dependencies or To add middleware
*/

type todoPostRequest struct {
	ID        int       `json:"id"`
	Data      string    `json:"data"`
	IsDone    bool      `json:"isDone"`
	CreatedAt time.Time `json:"create_at"`
	IsStared  bool      `json:"isStared"`
}

func AddTodo(todos models.Adder, ch *amqp.Channel) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := todoPostRequest{}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//c.Bind(&requestBody)
		id := rand.Int()
		todo := models.Todo{
			ID:       id,
			Data:     requestBody.Data,
			IsStared: requestBody.IsStared,
		}

		fmt.Println(todo)

		todos.Add(todo)
		strTodo, _ := json.Marshal(todo)

		message := amqp.Publishing{
			Body: []byte(strTodo),
		}

		// We publish the message to the exahange we created earlier
		err := ch.Publish("", constants.QUEUE_NAME, false, false, message)

		if err != nil {
			logger.FailOnError(err, "error publishing a message to the queue")
		}

		c.Status(http.StatusNoContent)

	}
}
