package Routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1:=r.Group("/v1"){
		v1.GET("todo",Controllers.GetTodos)
		v1.POST("todo",Controllers.CreateATodo)
		v1.GET("todo/:id",Controllers.UpdateATodo)
		v1.DELETE("todo/:id",Controllers.DeleteATodo)
		vq.PUT("todo/:id",Controllers.update)

	}
	return r

}
