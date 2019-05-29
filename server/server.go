package server

import (
	user "sample-api-gin/controller"

	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run(":8080")
}

func router() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		userCtrl := new(user.Controller)
		users.GET("", userCtrl.GetAllUser)
		users.GET("/:id", userCtrl.GetUser)
		users.POST("", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}

	// todos := r.Group("/todos")
	// {
	// 	todoCtrl := todo.Controller{}
	// 	users.GET("", todoCtrl.Index)
	// 	users.GET("/:id", todoCtrl.show)
	// 	// users.POST("", ctrl.Create)
	// 	// users.PUT("/:id", ctrl.Update)
	// 	// users.DELETE("/:id", ctrl.Delete)
	// }

	return r
}
