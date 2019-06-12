package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/temp-go-dev/sample-api-gin/config"
	"github.com/temp-go-dev/sample-api-gin/model"
	"github.com/temp-go-dev/sample-api-gin/service"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v8"
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
	logger := config.GetLogger()
	u := model.User{}

	fmt.Println("create")
	if err := c.BindJSON(&u); err != nil {
		fmt.Println("vali Error")
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
		return
	}

	// タグ v-post にもとづき validation
	config := &validator.Config{TagName: "v-post"}
	validate := validator.New(config)
	if err := validate.Struct(u); err != nil {
		fmt.Println("v-post vali Error")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("errField:%s ", err.Field)
			fmt.Printf("errType:%s\n", err.Tag)
		}

		logger.Error("error", zap.Error(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

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
