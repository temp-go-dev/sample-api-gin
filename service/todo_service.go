package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/model"
)

// TodoService procides user's behavior
type TodoService struct{}

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

// GetAllTodoTran Transactを使用した実装
func (s TodoService) GetAllTodoTran(uid string) ([]model.Todo, error) {
	db := db.GetDB()
	todos := []model.Todo{}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {

		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓

		// SELECT実行
		err := tx.Raw("SELECT * FROM todo where user_id = ?", uid).Scan(&todos).Error
		// GORM Find を利用したSELECT
		// tx.Table("todo").Where("user_id = ?", uid).Find(&todos)
		if err != nil {
			return nil, err
		}
		return todos, nil
		//↑↑↑ トランザクション対象の処理を記載 ↑↑↑

	})
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

// CreateTodosTran Transactを使用した実装
func (s TodoService) CreateTodosTran(todos model.Todos) ([]string, error) {
	db := db.GetDB()
	todoID := []string{}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {

		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓
		for _, todo := range todos.Todo {
			uuid := uuid.New()
			uuidStr := uuid.String()
			todo.ID = uuidStr
			errEvent := CreateTodo(tx, todo)
			todoID = append(todoID, uuidStr)
			if errEvent != nil {
				return nil, errEvent
			}
		}
		return todoID, nil
		//↑↑↑ トランザクション対象の処理を記載 ↑↑↑

	})
	if err != nil {
		return nil, err
	}
	return todoID, nil
}

// insert001 登録001
func insert001(db *gorm.DB, todos model.Todos) ([]string, error) {
	todoID := []string{}
	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := TransactNest(db, false, func(tx *gorm.DB) (interface{}, error) {
		for _, todo := range todos.Todo {
			uuid := uuid.New()
			uuidStr := uuid.String()
			todo.ID = uuidStr
			errEvent := CreateTodo(tx, todo)
			todoID = append(todoID, uuidStr)
			if errEvent != nil {
				return nil, errEvent
			}
		}
		return todoID, nil
	})
	if err != nil {
		return nil, err
	}
	return todoID, nil
}

// CreateTodosTranNest TransactNestを使用した実装
func (s TodoService) CreateTodosTranNest(todos model.Todos) ([]string, error) {
	fmt.Println("CreateTodosTranNest")
	db := db.GetDB().Begin()
	todoID := []string{}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := TransactNest(db, true, func(tx *gorm.DB) (interface{}, error) {
		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓

		// insert001 ネストしたトランザクション処理 Begin済みのDBを渡す
		todoID, _ = insert001(tx, todos)
		for i, todo := range todos.Todo {
			fmt.Println(i)
			if i > 0 {
				// 強制的にerrorにする
				panic("panic!!")
			}
			uuid := uuid.New()
			uuidStr := uuid.String()
			todo.ID = uuidStr
			errEvent := CreateTodo(tx, todo)
			todoID = append(todoID, uuidStr)
			if errEvent != nil {
				return nil, errEvent
			}
		}
		return todoID, nil
		//↑↑↑ トランザクション対象の処理を記載 ↑↑↑
	})
	if err != nil {
		return nil, err
	}
	return todoID, nil
}

// CreateTodo todoのINSERT
func CreateTodo(db *gorm.DB, todo model.Todo) error {
	err := db.Exec("INSERT INTO sampledb.todo VALUES (?,?,?,?,?,?,?,?,?,?);", todo.ID, todo.UserID, todo.Title, todo.Contents, todo.Start, todo.Due, todo.ActualStart, todo.ActualEnd, todo.Status, todo.Version).Error
	// err := db.Table("todo").Create(&todo).Error
	if err != nil {
		fmt.Println("DB error")
		return err
	}
	return nil
}

// Transact トランザクション実行
func Transact(db *gorm.DB, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	// 無名関数にBeginしたDBを渡して実行する
	data, err = txFunc(tx)
	return
}

// TransactNest トランザクション実行ネスト
func TransactNest(tx *gorm.DB, commit bool, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Rollback recover")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("Rollback err")
			tx.Rollback()
		} else if commit == true {
			fmt.Println("commit")
			err = tx.Commit().Error
		}
		fmt.Println("commitしない")
	}()
	// 無名関数にBeginしたDBを渡して実行する
	data, err = txFunc(tx)
	return
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
