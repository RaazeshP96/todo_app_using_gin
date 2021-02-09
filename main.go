package main

import (
	"fmt"

	"github.com/RaazeshP96/todo_app_using_gin/Config"
	"github.com/RaazeshP96/todo_app_using_gin/Models"
	"github.com/RaazeshP96/todo_app_using_gin/Routes"
	"gorm.io/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
