package user

import (
	"fmt"

	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/model"
)

// Service procides user's behavior
type Service struct{}

// User is alias of model.User struct
type User model.User

// GetAllUser is get all User
func (s Service) GetAllUser() ([]model.User, error) {
	db := db.GetDB()
	users := []model.User{}

	// SELECT実行
	err := db.Table("user").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser is get all User
func (s Service) GetUser(id string) ([]model.User, error) {
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
func (s Service) CreateUser(user model.User) (string, error) {
	db := db.GetDB()

	// Create実行
	// err := db.Raw("SELECT * FROM user where id = ?", id).Scan(&users).Error
	err := db.Table("user").Create(&user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

// UpdateUser ユーザを更新
func (s Service) UpdateUser(user model.User) (string, error) {
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
func (s Service) DeleteUser(id string) (string, error) {
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
