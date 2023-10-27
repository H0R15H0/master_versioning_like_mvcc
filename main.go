package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// データベースの情報
	DBMS = "mysql"

	// ユーザー名
	DMUSER = "root"

	// パスワード
	DBPASS = "mysql"

	// ホスト
	DBHOST = "localhost"

	// ポート
	DBPORT = "1234"

	// groupの数
	GROUP = 100

	// item codeの数
	ITEMCODE = 3000

	// item のバージョン数
	ITEMVERSION = 3
)

// Item テーブルの構造体
type Item struct {
	ID        int
	GroupID   int
	Code      string
	Name      string
	CreatedAt time.Time
}

func main() {
	// DBに接続
	db, err := sql.Open(DBMS, DMUSER+":"+DBPASS+"@tcp("+DBHOST+":"+DBPORT+")/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データの作成
	items := make([]Item, GROUP*ITEMCODE*ITEMVERSION)
	// 2000/01/01でTimeを初期化
	t, _ := time.Parse("2006/01/02", "2000/01/01")
	for i := 0; i < GROUP; i++ {
		for j := 0; j < ITEMCODE; j++ {
			for k := 0; k < ITEMVERSION; k++ {
				items[i*ITEMCODE*ITEMVERSION+j*ITEMVERSION+k] = Item{
					GroupID:   i,
					Code:      strconv.Itoa(j),
					Name:      strconv.Itoa(i) + "name" + strconv.Itoa(j) + "v" + strconv.Itoa(k),
					CreatedAt: t.AddDate(0, 0, k),
				}
			}
		}
	}

	start := time.Now()
	// データの挿入をバルクインサートで行う
	// データを1000個ずつに分割する
	for i := 0; i < len(items); i += 1000 {
		end := i + 1000
		if end > len(items) {
			end = len(items)
		}
		query := "INSERT INTO items (group_id, code, name, created_at) VALUES "
		for _, item := range items[i:end] {
			query += "(" + strconv.Itoa(item.GroupID) + ", '" + item.Code + "', '" + item.Name + "', '" + item.CreatedAt.Format("2006-01-02 15:04:05") + "'),"
		}
		// 最後のカンマを削除
		query = query[:len(query)-1]
		// バルクインサートの実行
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	end := time.Now()
	fmt.Printf("処理時間: %f秒\n", (end.Sub(start)).Seconds())
}
