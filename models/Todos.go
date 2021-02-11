package models

import (
	"fmt"

	Config "github.com/RaazeshP96/todo_app_using_gin/config"
)

func GetAllTodos(todo *[]Todo) (err error) {
	if err = Config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateATodo(todo *Todo) (err error) {
	if err = Config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetATodo(todo *Todo, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	Config.DB.Save(todo)
	return nil
}

func DeleteATodo(todo *Todo, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(todo)
	return nil
}
