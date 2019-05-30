package controller

import (
	"fmt"
	"sample-api-gin/model"
	"sample-api-gin/service"

	"github.com/gin-gonic/gin"
)

// UserController is user controlller
type UserController struct{}

// GetAllUser action: GET /users
func (pc UserController) GetAllUser(c *gin.Context) {
	var s service.UserService
	p, err := s.GetAllUser()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// GetUser action: GET /users/id
func (pc UserController) GetUser(c *gin.Context) {
	id := c.Param("id")

	var s service.UserService
	p, err := s.GetUser(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: Create /users
func (pc UserController) Create(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)

	var s service.UserService
	p, err := s.CreateUser(u)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: UPDATE /users/id
func (pc UserController) Update(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)

	var s service.UserService
	p, err := s.UpdateUser(u)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/id
func (pc UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	var s service.UserService
	p, err := s.DeleteUser(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
