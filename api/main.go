// api/main.go
package main

import (
	"fmt"
	"github.com/couchbase/gocb/v2"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Couchbase接続設定
	cluster, err := gocb.Connect("couchbase://localhost", gocb.ClusterOptions{
		Username: "admin",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("Failed to connect to Couchbase: %v\n", err)
	}

	// バケット名とコレクション名を指定
	bucket := cluster.Bucket("default")
	collection := bucket.DefaultCollection()

	// Echoのインスタンス作成
	e := echo.New()

	// APIエンドポイント作成
	e.GET("/get-data", func(c echo.Context) error {
		// Couchbaseからデータ取得
		var result map[string]interface{}
		getResult, err := collection.Get("document_id", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
		}

		// ドキュメントの結果をパース
		err = getResult.Content(&result)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error parsing data: %v", err))
		}

		// データを返却
		return c.JSON(http.StatusOK, result)
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
