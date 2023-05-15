package model

import (
	"strconv"
	"time"

	"github.com/rsingla/todo-service/entity"
)

// Todo is a struct that represents a todo item
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DateCreated time.Time `json:"dateCreated"`
	DateUpdated time.Time `json:"dateUpdated"`
}

// ConvertToTodo converts a TodoEntity struct to a Todo struct
func ConvertToTodo(todoE entity.TodoEntity) Todo {
	return Todo{
		ID:          strconv.Itoa(int(todoE.ID)),
		Title:       todoE.Title,
		Description: todoE.Description,
		Completed:   todoE.Completed,
		DateCreated: todoE.DateCreated,
		DateUpdated: todoE.DateUpdated,
	}
}
