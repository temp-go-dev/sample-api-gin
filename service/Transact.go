package service

import (	
	"github.com/jinzhu/gorm"
)

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
