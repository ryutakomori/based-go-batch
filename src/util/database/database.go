package Util

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":3306)"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/?charset=utf8mb4"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

// クエリ実行
func Exec(sql string) {
	db := Connect()

	tx := db.Begin()

	if err := db.Exec(sql).Error; err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tx.Commit()
	defer db.Close()
}
