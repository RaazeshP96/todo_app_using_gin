package routes

import (
	"github.com/RaazeshP96/todo_app_using_gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)

	}
	return r

}
