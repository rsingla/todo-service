package mydb

import (
	"github.com/rsingla/todo-service/entity"
	"gorm.io/gorm"
)

type TodoDBInterface struct {
	db *gorm.DB
}

func NewTodoDBInterface(db *gorm.DB) *TodoDBInterface {
	return &TodoDBInterface{
		db: db,
	}
}

func (tdb *TodoDBInterface) Create(todo *entity.TodoEntity) error {
	return tdb.db.Create(todo).Error
}

func (tdb *TodoDBInterface) FindAll() ([]entity.TodoEntity, error) {
	var todos []entity.TodoEntity
	err := tdb.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tdb *TodoDBInterface) FindByID(id uint) (*entity.TodoEntity, error) {
	var todo entity.TodoEntity
	err := tdb.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (tdb *TodoDBInterface) Update(todo *entity.TodoEntity) error {
	return tdb.db.Save(todo).Error
}

func (tdb *TodoDBInterface) Delete(todo *entity.TodoEntity) error {
	return tdb.db.Delete(todo).Error
}
