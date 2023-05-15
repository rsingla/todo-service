package entity

import "time"

// Todo is a struct that represents a todo item
type TodoEntity struct {
	ID          uint      `gorm:"primary_key"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Completed   bool      `gorm:"not null"`
	DateCreated time.Time `gorm:"not null"`
	DateUpdated time.Time `gorm:"not null"`
}

// ConvertToEntity converts a Todo struct to TodoEntity
func ConvertToPostEntity(todo model.Todo) TodoEntity {
	return TodoEntity{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		DateCreated: todo.DateCreated,
		DateUpdated: todo.DateUpdated,
	}
}
