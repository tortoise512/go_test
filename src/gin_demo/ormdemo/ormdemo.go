package main

import (
	"fmt"
	// "strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义一个雇员
type Emp struct {
	Id         int
	Name       string
	Age        int
	Gender     []uint8    // 对应数据库中的bit
	Birthday   *time.Time // 对应数据库中的date
	Salary     float32
	Department string
}

// 定义一个学生
type Student struct {
	Name string
	Age  int
}

// 实例化结构体的时候，time.Time输入字符串报错，所以这里定义了一个将字符串转换成时间的函数
func s2t(s string) time.Time {
	t, _ := time.Parse("2006-01-02", s)
	return t
}

func main() {
	dsn := "root:123456.com@(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.Table("employes").AutoMigrate(&Emp{})

}
