package server

import (
	"fmt"

	// ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/temp-go-dev/sample-api-gin/controller"
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
		todos.GET("/:id", sampleMiddleware("start: todo get", "end  : todo get"), todoCtrl.GetAllTodo)
		// 	users.GET("/:id", todoCtrl.show)
		todos.POST("", sampleMiddleware("start: todo POST", "end  : todo POST"), todoCtrl.Create)
		todos.POST("/error", sampleMiddleware("start: todo POST", "end  : todo POST"), todoCtrl.Create1)
		// 	// users.PUT("/:id", ctrl.Update)
		// 	// users.DELETE("/:id", ctrl.Delete)
	}
	// ginのログをzapのログに出す　うまくいってない
	// logger := config.GetLogger()
	// r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(logger, true))

	return r
}

// sampleMiddleware middlewareテスト
func sampleMiddleware(start string, end string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(start)
		c.Next()
		fmt.Println(end)
	}
}
