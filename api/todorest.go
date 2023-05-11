package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rsingla/todo-service/model"
)

var todoList []model.Todo

func init() {
	todoList = []model.Todo{
		{
			ID:          "1",
			Title:       "Go to the grocery store",
			Description: "Buy milk, eggs, bread, and Diet Coke",
			Completed:   false,
			CreatedDate: "2021-01-01",
			UpdatedDate: "2021-01-01",
		},
		{
			ID:          "2",
			Title:       "Go to the hardware store",
			Description: "Buy a hammer",
			Completed:   false,
			CreatedDate: "2023-05-10",
			UpdatedDate: "2021-05-10",
		},
		{
			ID:          "3",
			Title:       "Go to the liquor store",
			Description: "Buy a bottle of wine",
			Completed:   false,
			CreatedDate: "2023-05-10",
			UpdatedDate: "2021-05-10",
		},
	}
}

// GetTodos returns a list of all todos
func GetTodos(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"body": todoList,
	})
}

// GetTodo returns a single todo
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todoList {
		if todo.ID == id {
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

	if todo.CreatedDate == "" {
		todo.CreatedDate = time.Now().Format("2006-01-02")
	}

	if todo.UpdatedDate == "" {
		todo.UpdatedDate = time.Now().Format("2006-01-02")
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

	for index, item := range todoList {
		if item.ID == id {
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
