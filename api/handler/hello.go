package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Name":    "Todo ap",
		"version": "0.01",
	})
}
