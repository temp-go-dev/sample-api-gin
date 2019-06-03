package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/temp-go-dev/sample-api-gin/config"
	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/server"
	user "github.com/temp-go-dev/sample-api-gin/controller"
	"github.com/gin-gonic/gin"
)

// Init aa
// func Init() {
// 	db := gormConnect()
// }

// func gormConnect() *gorm.DB {
// 	DBMS := "mysql"
// 	USER := "user"
// 	PASS := "password"
// 	PROTOCOL := "tcp(localhost:3306)"
// 	DBNAME := "sampledb"
// 	// parseTime=trueを指定しないとdatetime→time.Timeへの変更でエラーが発生する。
// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
// 	db, err := gorm.Open(DBMS, CONNECT)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

func main() {
	// db := gormConnect()
	// defer db.Close()
	// db.LogMode(true)

	// r := gin.Default()

	// r.GET("/hello", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello world")
	// })

	// // Get user list.
	// r.GET("/users", func(c *gin.Context) {
	// 	users := []model.User{}

	// 	// SELECT実行
	// 	db.Table("user").Find(&users)
	// 	// db.Raw("SELECT * FROM todo").Scan(&todos)
	// 	c.JSON(http.StatusOK, users)
	// })

	// // Create user.
	// r.POST("/users", func(c *gin.Context) {
	// 	fmt.Println("create user")

	// 	user := model.User{}
	// 	c.BindJSON(&user)

	// 	// INSERT実行
	// 	db.Table("user").Create(&user)

	// 	c.JSON(http.StatusOK, user.ID)
	// })

	// // Update user.
	// r.PUT("/users/:id", func(c *gin.Context) {
	// 	fmt.Println("update user")

	// 	id := c.Param("id")
	// 	user := model.User{}

	// 	c.BindJSON(&user)
	// 	fmt.Println(id)
	// 	fmt.Println(user)

	// 	// UPDATE実行
	// 	db.Table("user").Save(&user)

	// 	c.JSON(http.StatusOK, user.ID)
	// })

	// // Delete user.
	// r.DELETE("/users/:id", func(c *gin.Context) {
	// 	fmt.Println("delete user")

	// 	id := c.Param("id")
	// 	user := model.User{}

	// 	// DELETE実行
	// 	db.Table("user").Where("id = ?", id).Delete(&user)
	// 	c.JSON(http.StatusOK, user.ID)
	// })

	// r.GET("/todos", func(c *gin.Context) {
	// 	todos := []model.Todo{}
	// 	db.Table("todo").Find(&todos)
	// 	// db.Raw("SELECT * FROM todo").Scan(&todos)
	// 	c.JSON(http.StatusOK, todos)
	// })

	// r.GET("/todos/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	todos := []model.Todo{}
	// 	db.Raw("SELECT * FROM todo where id = ?", id).Scan(&todos)
	// 	c.JSON(http.StatusOK, todos)
	// })

	// r.Run(":8080")

	config.Init()
	fmt.Println(config.GetProperties())

	db.Init()
	//server.Init()
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
