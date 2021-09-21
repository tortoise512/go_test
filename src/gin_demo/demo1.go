package main

import (
	"database/sql/driver"
	"fmt"
	// "strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type user struct {
	Id         int
	Name       string // 首字母大写，不然后面的Get拿不到值会报错，这个字段需要跟数据库字段对应
	Age        int
	Gender     []uint8 // go语言目前没有跟bit类型对应的数据类型
	Birthday   time.Time
	Salary     float32
	Department string
}

type student struct {
	Name string
	Age  int
}

func (u student) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

func main() {
	// 连接数据库
	var db *sqlx.DB
	dsn := "root:123456.com@(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=true"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	// 批量插入
	// 方式一：手动拼接
	u1 := student{Name: "七米", Age: 18}
	u2 := student{Name: "q1mi", Age: 28}
	u3 := student{Name: "小王子", Age: 38}
	users := []interface{}{u1, u2, u3}

	// 方式二 sqlx.In
	query, args, _ := sqlx.In(
		"INSERT INTO stu (name, age) VALUES (?), (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err1 := db.Exec(query, args...)
	fmt.Println(err1)
	// 方式三 NamedExec实现批量插入

}
