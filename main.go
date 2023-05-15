package main

import (
	"net/http"

	"github.com/rsingla/todo-service/api"

	"github.com/gin-gonic/gin"

	"github.com/rsingla/todo-service/mydb"
)

func main() {

	mydb.Connect()

	const PORT = "8081"

	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/health", health)

	r.GET("/todos", api.GetTodos)
	r.GET("/todo/:id", api.GetTodo)
	r.POST("/todo", api.CreateTodo)
	r.PUT("/todo/:id", api.UpdateTodo)
	r.DELETE("/todo/:id", api.DeleteTodo)

	r.Run(":" + PORT)

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Service is healthy!",
	})
}
