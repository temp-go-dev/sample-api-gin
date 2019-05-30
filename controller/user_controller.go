package user

import (
	"fmt"
	"github.com/temp-go-dev/sample-api-gin/model"
	user "github.com/temp-go-dev/sample-api-gin/service"

	"github.com/gin-gonic/gin"
)

// Controller is user controlller
type Controller struct{}

// GetAllUser action: GET /users
func (pc Controller) GetAllUser(c *gin.Context) {
	var s user.Service
	p, err := s.GetAllUser()

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// GetUser action: GET /users/id
func (pc Controller) GetUser(c *gin.Context) {
	id := c.Param("id")

	var s user.Service
	p, err := s.GetUser(id)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: Create /users
func (pc Controller) Create(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)

	var s user.Service
	p, err := s.CreateUser(u)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: UPDATE /users/id
func (pc Controller) Update(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)

	var s user.Service
	p, err := s.UpdateUser(u)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /users/id
func (pc Controller) Delete(c *gin.Context) {
	id := c.Param("id")

	var s user.Service
	p, err := s.DeleteUser(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}
}
