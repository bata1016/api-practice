package server

import (
	user "github.com/bata1016/api-practice/controller"
	"github.com/gin-gonic/gin"
)

// this is initialize server
func Init() {
	router := router()
	router.Run()
}

func router() *gin.Engine {
	router := gin.Default()

	userPath := router.Group("/users")
	{
		ctrl := user.Controller{}
		userPath.GET("", ctrl.Index)
		userPath.GET("/:id", ctrl.Show)
		userPath.POST("/:id", ctrl.Create)
		userPath.DELETE("/:id", ctrl.Delete)
	}
	return router
}
