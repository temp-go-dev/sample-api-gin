package main


import (
	// ブランクインポートはメイン関数で行えとgolintに怒られるため、main関数に記載
    // ブランクインポートはグローバルレベルなのでここでどこでimportしてもみえる。
    // グローバルだからサブパッケージでやらずに、mainでやれとのこと
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/temp-go-dev/sample-api-gin/config"
	"github.com/temp-go-dev/sample-api-gin/db"
	"github.com/temp-go-dev/sample-api-gin/server"
)

func main() {
	config.Init()
	db.Init()
	server.Init()
}
