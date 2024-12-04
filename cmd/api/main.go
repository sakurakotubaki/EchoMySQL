package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"myapi/internal/domain/model"
	"myapi/internal/infrastructure/mysql"
	"myapi/internal/interface/handler"
	"myapi/internal/usecase"
)

func main() {
	// データベース接続
	db, err := gorm.Open("mysql", "jboy:1234qw@tcp(localhost:3306)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// マイグレーション
	db.AutoMigrate(&model.ShoppingItem{})

	// リポジトリの初期化
	repo := mysql.NewShoppingRepository(db)

	// ユースケースの初期化
	uc := usecase.NewShoppingUsecase(repo)

	// ハンドラーの初期化
	h := handler.NewShoppingHandler(uc)

	// Echoの初期化
	e := echo.New()

	// ルーティング
	e.POST("/shopping", h.CreateItem)
	e.GET("/shopping", h.GetAllItems)
	e.GET("/shopping/:id", h.GetItem)
	e.PUT("/shopping/:id", h.UpdateItem)
	e.DELETE("/shopping/:id", h.DeleteItem)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
