package db

import (
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Init DB初期化
func Init() {
	DBMS := "mysql"
	USER := "user"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "sampledb"
	// parseTime=trueを指定しないとdatetime→time.Timeへの変更でエラーが発生する。
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(db)
	db.LogMode(true)

}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
