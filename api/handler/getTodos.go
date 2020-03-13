package handler

import (
	//"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/razzkumar/todo/api/models"
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

func GetTodos(todo models.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos := todo.GetAll()
		c.JSON(http.StatusOK, todos)
	}
}
