package service

import (
	"fmt"
	"sample-api-gin/db"
	"sample-api-gin/model"
)

// TodoService procides user's behavior
type TodoService struct{}

// User is alias of model.User struct
// type User model.User

// GetAllTodo ユーザIDに紐づくTODOを取得
func (s TodoService) GetAllTodo(uid string) ([]model.Todo, error) {
	db := db.GetDB()
	todos := []model.Todo{}

	// SELECT実行
	err := db.Raw("SELECT * FROM todo where user_id = ?", uid).Scan(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// // GetUser is get all User
// func (s Service) GetUser(id string) ([]model.User, error) {
// 	db := db.GetDB()
// 	users := []model.User{}

// 	// SELECT実行
// 	err := db.Raw("SELECT * FROM user where id = ?", id).Scan(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// CreateTodos Todoの作成
func (s TodoService) CreateTodos(todos model.Todos) (string, error) {
	db := db.GetDB()

	// //トランザクション開始
	tx := db.Begin()
	if tx.Error != nil {
		return "", tx.Error
	}

	// Create実行
	for _, todo := range todos.Todo {
		err := tx.Table("todo").Create(&todo).Error
		if err != nil {
			fmt.Println("DB error")
			// ロールバックして終了
			tx.Rollback()
			return "", err
		}
	}
	// コミットして終了
	tx.Commit()
	return "", nil
}

// // UpdateUser ユーザを更新
// func (s Service) UpdateUser(user model.User) (string, error) {
// 	fmt.Print("update")
// 	db := db.GetDB()
// 	// user := model.User{}

// 	// UPDATE
// 	err := db.Table("user").Save(&user).Error
// 	if err != nil {
// 		return "", err
// 	}
// 	return user.ID, nil
// }

// // DeleteUser ユーザを削除
// func (s Service) DeleteUser(id string) (string, error) {
// 	fmt.Print("delete")
// 	db := db.GetDB()
// 	user := model.User{}

// 	// DELETE実行 存在チェック後、存在した場合は削除実行
// 	err := db.Table("user").Where("id = ?", id).Delete(&user).Error
// 	if err != nil {
// 		return "", err
// 	}
// 	return user.ID, nil
// }
