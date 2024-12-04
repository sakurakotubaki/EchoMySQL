package main

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

// ShoppingItemは、買い物リストの構造体です。
type ShoppingItem struct {
	gorm.Model
	Name string `json:"name"`
}

// DBは、データベースのインスタンスです。
var DB *gorm.DB

// initは、データベースの初期化を行います。
func init() {
	// DBの初期化
	db, err := gorm.Open("mysql", "jboy:1234qw@tcp(localhost:3306)/myapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// エラー処理
		panic(err.Error()) // エラーメッセージを表示するように修正
	}
	// マイグレーション
	db.AutoMigrate(&ShoppingItem{})
	DB = db
}

func main() {
	// Echoのインスタンス作成
	e := echo.New()

	// カスタムエラーハンドラの設定
	e.HTTPErrorHandler = customHTTPErrorHandler

	// ミドルウェアの設定
	e.POST("/shopping", createItem)
	e.GET("/shopping", getAllItems) // Add this line
	e.GET("/shopping/:id", getItem)
	e.PUT("/shopping/:id", updateItem)
	e.DELETE("/shopping/:id", deleteItem)
	// /helloのURLにアクセスしたら、Hello GolangというHTMLを返す
	e.GET("/hello", getGreet)
	// サーバー起動
	e.Start(":8080")
}

func getGreet(c echo.Context) error {
	return c.File("index.html")
}

// MySQLにデータを追加
func createItem(c echo.Context) error {
	item := new(ShoppingItem)
	if err := c.Bind(item); err != nil {
		return err
	}
	DB.Create(&item)
	return c.JSON(http.StatusCreated, item)
}

// MySQLからデータを全て取得
func getAllItems(c echo.Context) error {
	items := []ShoppingItem{}
	DB.Find(&items)
	return c.JSON(http.StatusOK, items)
}

// MySQLからデータを1つ取得
func getItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item := new(ShoppingItem)
	if DB.First(&item, id).RecordNotFound() {
		return echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}
	return c.JSON(http.StatusOK, item)
}

// MySQLのデータを更新
func updateItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	newItem := new(ShoppingItem)
	if err := c.Bind(newItem); err != nil {
		return err
	}
	oldItem := new(ShoppingItem)
	DB.First(oldItem, id)
	oldItem.Name = newItem.Name
	DB.Save(oldItem)
	return c.JSON(http.StatusOK, oldItem)
}

// MySQLのデータを削除
func deleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item := new(ShoppingItem)
	DB.First(&item, id)
	DB.Delete(item)
	return c.NoContent(http.StatusNoContent)
}

// カスタムエラーハンドラ
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := http.StatusText(code)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	if code == http.StatusNotFound {
		// 404エラーの場合は404.htmlを返す
		c.File("404.html")
		return
	}

	c.JSON(code, map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	})
}
