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
	// serviceの呼び出し
	var s service.UserService
	p, err := s.GetAllUser()
	if err != nil { // エラーが設定されている場合
		logger := config.GetLogger()
		switch e := err.(type) {
		case *service.ErrorMessage: // ErrorMessageが設定されている場合、
			// ログ出力
			logger.Error("error", zap.Error(e.Detail))
			// レスポンスボディを生成して終了
			c.JSON(e.StatusCd, e)
			return
		default: // されていない場合サーバエラーを返却
			logger.Error("error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	} else {
		c.JSON(http.StatusOK, p)
		return
	}
}

// GetUser action: GET /users/id
func (pc UserController) GetUser(c *gin.Context) {
	// パラメータ取得
	id := c.Param("id")

	// serviceの呼び出し
	var s service.UserService
	p, err := s.GetUser(id)
	if err != nil { // エラーが設定されている場合
		logger := config.GetLogger()
		switch e := err.(type) {
		case *service.ErrorMessage: // ErrorMessageが設定されている場合、
			// ログ出力
			logger.Error("error", zap.Error(e.Detail))
			// レスポンスボディを生成して終了
			c.JSON(e.StatusCd, e)
			return
		default: // されていない場合サーバエラーを返却
			logger.Error("error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	} else {
		c.JSON(http.StatusOK, p)
		return
	}
}

// Create action: Create /users
func (pc UserController) Create(c *gin.Context) {
	logger := config.GetLogger()
	u := model.User{}

	if err := c.BindJSON(&u); err != nil {
		// バインドエラーの場合、400で終了
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// タグ v-post にもとづき validation
	config := &validator.Config{TagName: "v-post"}
	validate := validator.New(config)
	if err := validate.Struct(u); err != nil {
		// validationエラーの場合
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("errField:%s ", err.Field)
			fmt.Printf("errType:%s\n", err.Tag)
		}
		// ログ出力
		logger.Error("error", zap.Error(err))
		// レスポンス生成して終了
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// serviceの呼び出し
	var s service.UserService
	p, err := s.CreateUser(u)
	if err != nil { // エラーが設定されている場合
		switch e := err.(type) {
		case *service.ErrorMessage: // ErrorMessageが設定されている場合、
			// ログ出力
			logger.Error("error", zap.Error(e.Detail))
			// レスポンスボディを生成して終了
			c.JSON(e.StatusCd, e)
			return
		default: // されていない場合サーバエラーを返却
			logger.Error("error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	} else {
		c.JSON(http.StatusOK, p)
		return
	}
}

// Update action: UPDATE /users/id
func (pc UserController) Update(c *gin.Context) {
	u := model.User{}

	// リクエストボディのバインド処理
	if err := c.BindJSON(&u); err != nil {
		// バインドエラーの場合、400で終了
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// serviceの呼び出し
	var s service.UserService
	p, err := s.UpdateUser(u)
	if err != nil { // エラーが設定されている場合
		logger := config.GetLogger()
		switch e := err.(type) {
		case *service.ErrorMessage: // ErrorMessageが設定されている場合、
			// ログ出力
			logger.Error("error", zap.Error(e.Detail))
			// レスポンスボディを生成して終了
			c.JSON(e.StatusCd, e)
			return
		default: // されていない場合サーバエラーを返却
			logger.Error("error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Delete action: DELETE /users/id
func (pc UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	// serviceの呼び出し
	var s service.UserService
	p, err := s.DeleteUser(id)
	if err != nil { // エラーが設定されている場合
		logger := config.GetLogger()
		switch e := err.(type) {
		case *service.ErrorMessage: // ErrorMessageが設定されている場合、
			// ログ出力
			logger.Error("error", zap.Error(e.Detail))
			// レスポンスボディを生成して終了
			c.JSON(e.StatusCd, e)
			return
		default: // されていない場合サーバエラーを返却
			logger.Error("error", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
			return
		}
	} else {
		c.JSON(http.StatusOK, p)
	}
}
