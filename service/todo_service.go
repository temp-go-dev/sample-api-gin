package service

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/model"
)

// TodoService procides user's behavior
type TodoService struct{}

// GetAllTodoTran ユーザIDに紐づくTODOの取得
func (s TodoService) GetAllTodoTran(uid string) ([]model.Todo, error) {
	db := db.GetDB()
	todos := []model.Todo{}

	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {
		// SELECT実行
		err := tx.Raw("SELECT * FROM todo where user_id = ?", uid).Scan(&todos).Error
		if err != nil {
			return nil, err
		}
		return todos, nil
	})
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func inserts(tx *gorm.DB, todos model.Todos) ([]string, error) {
	todoID := []string{}
	for _, todo := range todos.Todo {
		uuid := uuid.New().String()
		todo.ID = uuid
		errEvent := CreateTodo(tx, todo)
		todoID = append(todoID, uuid)
		if errEvent != nil {
			return nil, errEvent
		}
	}
	return todoID, nil
}

// CreateTodosTran Todoを生成する
func (s TodoService) CreateTodosTran(todos model.Todos) ([]string, error) {
	db := db.GetDB()
	todoID := []string{}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := Transact(db, func(tx *gorm.DB) (interface{}, error) {
		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓
		for _, todo := range todos.Todo {
			uuid := uuid.New().String()
			todo.ID = uuid
			errEvent := CreateTodo(tx, todo)
			todoID = append(todoID, uuid)
			if errEvent != nil {
				return nil, &DbError{"dbError", errEvent}
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
	if err != nil {
		return err
	}
	return nil
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
	db := db.GetDB().Begin()
	todoID := []string{}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := TransactNest(db, true, func(tx *gorm.DB) (interface{}, error) {
		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓

		// insert001 ネストしたトランザクション処理 Begin済みのDBを渡す
		todoID, _ = insert001(tx, todos)

		// 一意制約でエラーにする
		uuid := uuid.New()
		uuidStr := uuid.String()

		for _, todo := range todos.Todo {
			todo.ID = uuidStr
			errEvent := CreateTodo(tx, todo)
			if errEvent != nil {
				return nil, errors.Wrap(errEvent, "DBアクセス処理でエラー")
			}
			todoID = append(todoID, uuidStr)
		}
		return todoID, nil
		//↑↑↑ トランザクション対象の処理を記載 ↑↑↑
	})
	if err != nil {
		return nil, errors.Wrap(err, "トランザクション処理でエラー")
	}
	return todoID, nil
}

// CreateTodosErrorHandling エラーハンドリングのサンプル
func (s TodoService) CreateTodosErrorHandling(todos model.Todos) ([]string, error) {
	db := db.GetDB().Begin()
	todoID := []string{}

	if len := len(todos.Todo); len == 0 {
		// 0件の場合エラー
		return nil, &CheckError{"error 登録対象がありません。", "E1001"}
	}

	// Transactにトランザクションを行いたい処理を実装した無名関数を渡す
	_, err := TransactNest(db, true, func(tx *gorm.DB) (interface{}, error) {
		// ↓↓↓ トランザクション対象の処理を記載 ↓↓↓

		// insert001 ネストしたトランザクション処理 Begin済みのDBを渡す
		todoID, _ = insert001(tx, todos)

		// 一意制約でエラーにする
		uuid := uuid.New()
		uuidStr := uuid.String()

		for _, todo := range todos.Todo {
			todo.ID = uuidStr
			errEvent := CreateTodo(tx, todo)
			if errEvent != nil {
				return nil, errors.Wrap(errEvent, "dbError")
				// return nil, errors.Wrap(errEvent, &DbError{"Dberror"})
			}
			todoID = append(todoID, uuidStr)
		}
		return todoID, nil
		//↑↑↑ トランザクション対象の処理を記載 ↑↑↑
	})
	if err != nil {
		return nil, errors.Wrap(err, "トランザクション処理でエラー")
	}
	return todoID, nil
}

// TransactNest トランザクション実行ネスト
func TransactNest(tx *gorm.DB, commit bool, txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else if commit == true {
			err = tx.Commit().Error
		}
	}()
	// 無名関数にBeginしたDBを渡して実行する
	data, err = txFunc(tx)
	return
}

// GetAllTodo 【使用していない】トランザクションなしの実装例
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

// CreateTodos 【未使用】トランザクションなしの実装 Todoの作成
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
			// ロールバックして終了
			tx.Rollback()
			return "", err
		}
	}
	// コミットして終了
	tx.Commit()
	return "", nil
}
