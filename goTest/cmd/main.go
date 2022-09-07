package main

import (
	"fmt"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "net/http"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(121.196.223.94:3306)/gotest")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer db.Close() // 注意这行代码要写在上面err判断的下面
}

func main() {
	//r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//id, err := r.LastInsertId()
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//
	//fmt.Println("insert succ:", id)

	//var person []Person
	//err := Db.Select(&person, "select * from person")
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	//for _, p := range person {
	//	fmt.Println(p)
	//}

	//c, err := redis.Dial("tcp", "121.196.223.94:6379")
	//if err != nil {
	//	fmt.Println("conn redis failed,", err)
	//	return
	//}
	//
	//fmt.Println("redis conn success")
	//
	//defer c.Close()

}
