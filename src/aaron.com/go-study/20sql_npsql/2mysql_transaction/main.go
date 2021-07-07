package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type student struct {
	id   int64
	name string
	age  int64
}

// initDb 连接初始化
func initDb() (err error) {
	// 1.连接字符串
	dsn := "root:123123@tcp(127.0.0.1:3306)/go_sample"
	// 2.连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("dsn failed,err:", err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		fmt.Println("open db failed,err:", err)
		return
	}
	//db.SetMaxOpenConns(1000) // 设置与数据库建立连接的最大数目
	//db.SetMaxIdleConns(100)	// 设置连接池中的最大闲置连接数
	return
}

// 事务操作示例
func transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "update student set age=30 where id=?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update student set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init db failed,err:", err)
		return
	}
	transaction()
}
