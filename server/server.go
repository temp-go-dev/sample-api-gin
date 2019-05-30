package server

import (
	"sample-api-gin/controller"

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
		userCtrl := new(controller.UserController)
		users.GET("", userCtrl.GetAllUser)
		users.GET("/:id", userCtrl.GetUser)
		users.POST("", userCtrl.Create)
		users.PUT("/:id", userCtrl.Update)
		users.DELETE("/:id", userCtrl.Delete)
	}

	todos := r.Group("/todos")
	{
		todoCtrl := controller.TodoController{}
		todos.GET("/:id", todoCtrl.GetAllTodo)
		// 	users.GET("/:id", todoCtrl.show)
		todos.POST("", todoCtrl.Create)
		// 	// users.PUT("/:id", ctrl.Update)
		// 	// users.DELETE("/:id", ctrl.Delete)
	}

	return r
}
