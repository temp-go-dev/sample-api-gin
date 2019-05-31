package service

import (
	"fmt"

	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/model"
)

// UserService procides user's behavior
type UserService struct{}

// User is alias of model.User struct
// type User model.User

// GetAllUser is get all User
func (s UserService) GetAllUser() ([]model.User, error) {
	db := db.GetDB()
	users := []model.User{}

	// SELECT実行
	// SQL直書きはGetUserで実装
	err := db.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser is get all User
func (s UserService) GetUser(id string) ([]model.User, error) {
	db := db.GetDB()
	users := []model.User{}

	// SELECT実行
	err := db.Raw("SELECT * FROM user where id = ?", id).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser is get all User
func (s UserService) CreateUser(user model.User) (string, error) {
	db := db.GetDB()
	//トランザクション開始
	tx := db.Begin()
	if tx.Error != nil {
		return "", tx.Error
	}

	// Create実行
	err := db.Table("user").Create(&user).Error
	if err != nil {
		tx.Rollback()
		return "", err
	}
	// コミットして終了
	tx.Commit()
	return user.ID, nil
}

// UpdateUser ユーザを更新
func (s UserService) UpdateUser(user model.User) (string, error) {
	fmt.Print("update")
	db := db.GetDB()
	// user := model.User{}

	// UPDATE
	err := db.Table("user").Save(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

// DeleteUser ユーザを削除
func (s UserService) DeleteUser(id string) (string, error) {
	fmt.Print("delete")
	db := db.GetDB()
	user := model.User{}

	// DELETE実行 存在チェック後、存在した場合は削除実行
	err := db.Table("user").Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}
