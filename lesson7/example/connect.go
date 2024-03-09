package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义一个全局对象db
var db *sqlx.DB

func initDB() {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	fmt.Println("connecting to MySQL...")
	return
}

func main() {
	//初始化连接
	initDB()
}
