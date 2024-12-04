# EchoMySQL

## プロジェクト構成
```
EchoMySQL/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   └── model/
│   │       └── shopping.go
│   ├── usecase/
│   │   ├── shopping_usecase.go
│   │   └── shopping_usecase_interface.go
│   ├── interface/
│   │   ├── handler/
│   │   │   └── shopping_handler.go
│   │   └── repository/
│   │       ├── shopping_repository.go
│   │       └── shopping_repository_interface.go
│   └── infrastructure/
│       └── mysql/
│           └── shopping_repository.go
└── pkg/
    └── error/
        └── error.go
```

## 環境構築

### 1. MySQLコンテナの起動
```bash
# イメージのビルド
docker build -t mysql-custom .

# コンテナの起動
docker run -d -p 3306:3306 --name mysql-container mysql-custom
```

### 2. データベース接続情報
- ホスト: localhost
- ポート: 3306
- データベース: myapp
- ユーザー名: jboy
- パスワード: 1234qw

## アプリケーションの実行

### 1. アプリケーションの起動
```bash
# プロジェクトのルートディレクトリで実行
go run cmd/api/main.go
```

### 2. アプリケーションの停止
```bash
# 実行中のプロセスを確認
lsof -i :8080

# プロセスの停止（PIDは実際の値に置き換えてください）
kill -9 <PID>
```

## API エンドポイント

### 商品の追加
```
POST http://localhost:8080/shopping
Content-Type: application/json

{
    "name": "商品名"
}
```

### 全商品の取得
```
GET http://localhost:8080/shopping
```

### 特定商品の取得
```
GET http://localhost:8080/shopping/:id
```

### 商品の更新
```
PUT http://localhost:8080/shopping/:id
Content-Type: application/json

{
    "name": "新しい商品名"
}
```

### 商品の削除
```
DELETE http://localhost:8080/shopping/:id
```

## その他の操作

### MySQLコンテナの管理
```bash
# コンテナの停止
docker stop mysql-container

# コンテナの削除
docker rm mysql-container

# コンテナのログ確認
docker logs mysql-container
```

### MySQLへの接続
```bash
# rootユーザーとして接続
docker exec -it mysql-container mysql -u root -p
# パスワード: root1234

# 一般ユーザーとして接続
docker exec -it mysql-container mysql -u jboy -p
# パスワード: 1234qw
```

## 🖼️スクリーンショット
POST MANを使用して、HTTP通信をおこなっています。
<img src="./image/01.png" alt="01"></img>
<img src="./image/02.png" alt="02"></img>
<img src="./image/03.png" alt="03"></img>

## 全体のコード
```go
package main

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db, err := gorm.Open("mysql", "root:1234@/MyData?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// エラー処理
		panic("failed to connect database")
	}
	// マイグレーション
	db.AutoMigrate(&ShoppingItem{})
	DB = db
}

func main() {
	// Echoのインスタンス作成
	e := echo.New()
  // ミドルウェアの設定
	e.POST("/shopping", createItem)
	e.GET("/shopping", getAllItems) 
	e.GET("/shopping/:id", getItem)
	e.PUT("/shopping/:id", updateItem)
	e.DELETE("/shopping/:id", deleteItem)
  // サーバー起動
	e.Start(":8080")
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