package db

import (
	"sample-api-gin/config"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Init DB初期化
func Init() {
	prop := config.GetProperties()

	// parseTime=trueを指定しないとdatetime→time.Timeへの変更でエラーが発生する。
	CONNECT := prop.User + ":" + prop.Pass + "@" + prop.Protocol + "/" + prop.Dbname + "?parseTime=true"
	db, err = gorm.Open(prop.Dbms, CONNECT)

	if err != nil {
		//　err発生時の処理要検討
		panic(err.Error())
	}
	// DBデバッグログの出力設定
	db.LogMode(true)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
