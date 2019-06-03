package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var (
	prop Properties
	err  error
)

// Init 環境変数読み込みとマッピングした構造体の初期化
func Init() {
	// dockerで読み込むよう方法に変更。envファイルはリポジトリを分けてsubmodule化
	// ファイル読み込み
	// err = godotenv.Load("environment.env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	err = envconfig.Process("env", &prop)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Properties ENVプロパティマッピング用
type Properties struct { //必須チェックを行うタグ
	Dbms     string `required:"true"`
	User     string `required:"true"`
	Pass     string `required:"true"`
	Protocol string `required:"true"`
	Dbname   string `required:"true"`
}

// GetProperties get prop
func GetProperties() Properties {
	return prop
}
