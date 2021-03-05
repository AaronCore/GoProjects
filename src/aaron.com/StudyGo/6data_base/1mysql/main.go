package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id   int64
	name string
	age  int64
}

var db *sql.DB

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

// queryRow
func queryRow() {
	querySql := "select id,name,age from student where id=?;"
	s := student{}
	err := db.QueryRow(querySql, 1).Scan(&s.id, &s.name, &s.age)
	if err != nil {
		fmt.Println("scan failed,err:", err)
		return
	}
	fmt.Println(s)
}

// queryAll
func queryAll() {
	querySql := "select id,name,age from student;"
	rows, err := db.Query(querySql)
	if err != nil {
		fmt.Println("exec query failed,err:", err)
		return
	}
	//关闭rows释放持有的数据库链接
	defer rows.Close()
	var students []student
	for rows.Next() {
		s := student{}
		err = rows.Scan(&s.id, &s.name, &s.age)
		if err != nil {
			fmt.Println("scan failed,err:", err)
			return
		}
		students = append(students, s)
	}
	fmt.Println(students)
}

func add(name string, age int64) {
	addSql := "insert into student(name, age) values (?,?)"
	ret, err := db.Exec(addSql, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func update() {
	sqlStr := "update student set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 28, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func delete() {
	sqlStr := "delete from student where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Println("init db failed,err:", err)
		return
	}
	//queryRow()
	queryAll()
	//add("王五",22)
	//update()
	//delete()
}
