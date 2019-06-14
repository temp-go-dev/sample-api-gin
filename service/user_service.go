package service

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/model"
)

// UserService procides user's behavior
type UserService struct{}

// GetAllUser is get all User
func (s UserService) GetAllUser() ([]model.User, error) {
	db := db.GetDB()
	users := []model.User{}

	// SELECT実行
	// SQL直書きはGetUserで実装
	err := db.Table("user").Find(&users).Error
	if err != nil {
		return nil, &ErrorMessage{
			StatusCd: http.StatusInternalServerError,
			Message:  "",
			ErrorCd:  "1005",
			Detail:   "DBerror",
			err:      err,
		}
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
		return nil, &ErrorMessage{
			StatusCd: http.StatusInternalServerError,
			Message:  "",
			ErrorCd:  "1005",
			Detail:   "DBerror",
			err:      err,
		}
	}
	return users, nil
}

// CreateUser ユーザを登録する
func (s UserService) CreateUser(user model.User) (string, error) {
	db := db.GetDB()

	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {
		err := db.Table("user").Create(&user).Error
		if err != nil {
			return "", &ErrorMessage{
				StatusCd: http.StatusInternalServerError,
				Message:  "",
				ErrorCd:  "1005",
				Detail:   "DBerror",
				err:      err,
			}
		}
		return user.ID, nil
	})
	if err != nil {
		return "", &ErrorMessage{
			StatusCd: http.StatusInternalServerError,
			Message:  "",
			ErrorCd:  "1005",
			Detail:   "DBerror",
			err:      err,
		}
	}
	return user.ID, nil
}

// UpdateUser ユーザを更新
func (s UserService) UpdateUser(user model.User) (string, error) {
	db := db.GetDB()

	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {
		err := db.Table("user").Save(&user).Error
		if err != nil {
			return "", &ErrorMessage{
				StatusCd: http.StatusInternalServerError,
				Message:  "",
				ErrorCd:  "1005",
				Detail:   "DBerror",
				err:      err,
			}
		}
		return user.ID, nil
	})
	if err != nil {
		return "", &ErrorMessage{
			StatusCd: http.StatusInternalServerError,
			Message:  "",
			ErrorCd:  "1005",
			Detail:   "DBerror",
			err:      err,
		}
	}
	return user.ID, nil
}

// DeleteUser ユーザを削除
func (s UserService) DeleteUser(id string) (string, error) {
	db := db.GetDB()

	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {
		err := db.Raw("DELETE FROM user WHERE id = ?", id).Error
		if err != nil {
			return "", &ErrorMessage{
				StatusCd: http.StatusInternalServerError,
				Message:  "",
				ErrorCd:  "1005",
				Detail:   "DBerror",
				err:      err,
			}
		}
		return id, nil
	})
	if err != nil {
		return "", &ErrorMessage{
			StatusCd: http.StatusInternalServerError,
			Message:  "",
			ErrorCd:  "1005",
			Detail:   "DBerror",
			err:      err,
		}
	}
	return id, nil
}
