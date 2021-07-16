package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type student struct {
	Id   int64
	Name string
	Age  int64
}

var db *sqlx.DB

// initDb 初始化db
func initDb() (err error) {
	db, err = sqlx.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/go_sample")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	return err
}

// get
func get() {
	s := student{}
	err := db.Get(&s, "select id,name,age from student where id=?;", 1)
	if err != nil {
		fmt.Println("get failed,err:", err)
		return
	}
	fmt.Println(s)
}

// query
func query() {
	s := []student{}
	err := db.Select(&s, "select id,name,age from student where id>?;", 0)
	if err != nil {
		fmt.Println("select failed,err:", err)
		return
	}
	fmt.Println(s)
}

// add
func add() {
	r, err := db.Exec("insert into student(name, age) values (?,?);", "A1", "30")
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	fmt.Println(id)
}

// update
func update() {
	r, err := db.Exec("update student set age=? where id = ?;", 20, 1)
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	fmt.Println(row)
}

// delete
func delete() {
	r, err := db.Exec("delete from student where id = ?;", 1)
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("exec failed,err:", err)
		return
	}
	fmt.Println(row)
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init db failed,err:", err)
		return
	}
	//get()
	//query()
	//add()
	//update()
	//delete()
}
