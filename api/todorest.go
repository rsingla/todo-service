package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/todo-service/model"
	"github.com/rsingla/todo-service/mydb"
)

var todoDB mydb.TodoDBInterface

func init() {
	dbConn := mydb.Connect()
	todoDB = mydb.NewTodoDBInterface(dbConn)
}

// GetTodos returns a list of all todos
func GetTodos(c *gin.Context) {
	todoList, err := todoDB.FindAll()
	c.JSON(http.StatusOK, todoList)
}

// GetTodo returns a single todo
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todoEntity, err := todoDB.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
	}

	todo := model.ConvertToTodoEntity(*todoEntity)

	c.JSON(http.StatusOK, gin.H{
		"body": todo,
	})
	return

}

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var todo model.Todo

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Title is required",
		})
		return
	}

	idVal := len(todoList) + 1
	todo.ID = strconv.Itoa(idVal)

	if todo.DateCreated == (time.Time{}) {
		todo.DateCreated = time.Now().UTC()
	}

	if todo.DateUpdated == (time.Time{}) {
		todo.DateUpdated = time.Now().UTC()
	}

	todoList = append(todoList, todo)

	c.JSON(http.StatusCreated, gin.H{
		"body": todo,
	})
}

// UpdateTodo updates an existing todo
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo model.Todo

	err := c.ShouldBindJSON(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if todo.ID != id {
		todo.ID = id
	}

	if todo.DateUpdated == (time.Time{}) {
		todo.DateUpdated = time.Now().UTC()
	}

	for index, item := range todoList {
		if item.ID == id {
			todoDB := todoList[index]
			todo.DateCreated = todoDB.DateCreated
			todoList[index] = todo
			c.JSON(http.StatusOK, gin.H{
				"body": todo,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Todo not found",
	})
}

// DeleteTodo deletes an existing todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	for index, item := range todoList {
		if item.ID == id {
			todoList = append(todoList[:index], todoList[index+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Todo deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message": "Todo not found",
	})
}
