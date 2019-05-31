package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temp-go-dev/sample-api-gin/model"
	"github.com/temp-go-dev/sample-api-gin/service"
)

// TodoController is todo controller
type TodoController struct{}

// GetAllTodo action: GET /todos
func (pc TodoController) GetAllTodo(c *gin.Context) {
	id := c.Param("id")

	var s service.TodoService
	// p, err := s.GetAllTodo(id)
	p, err := s.GetAllTodoTran(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// // GetUser action: GET /users/id
// func (pc UserController) GetUser(c *gin.Context) {
// 	id := c.Param("id")

// 	var s service.UserService
// 	p, err := s.GetUser(id)

// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// Create action: Create /todos
func (pc TodoController) Create(c *gin.Context) {
	todos := model.Todos{}
	if err := c.BindJSON(&todos); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}
	var s service.TodoService
	p, err := s.CreateTodos(todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errMessage": err})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Create1 action: Create /todos
func (pc TodoController) Create1(c *gin.Context) {
	todos := model.Todos{}
	if err := c.BindJSON(&todos); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	for i, todo := range todos.Todo {
		fmt.Println("ループインデックス:", i)
		fmt.Println(todo)
	}

	var s service.TodoService
	p, err := s.CreateTodosTran(todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "errMessage": err})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// // Update action: UPDATE /users/id
// func (pc UserController) Update(c *gin.Context) {
// 	u := model.User{}
// 	c.BindJSON(&u)

// 	var s service.UserService
// 	p, err := s.UpdateUser(u)
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }

// // Delete action: DELETE /users/id
// func (pc UserController) Delete(c *gin.Context) {
// 	id := c.Param("id")

// 	var s service.UserService
// 	p, err := s.DeleteUser(id)
// 	if err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, p)
// 	}
// }
