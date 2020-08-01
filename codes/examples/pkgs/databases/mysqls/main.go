package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
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

var db *sqlx.DB

func init() {
	fmt.Println("hi, mysql test")
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	database, err := sqlx.Open("mysql", "root:333333@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql err -> ", err)
	}
	db = database
}

func insert(sql string) (insertId int64) {
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal("insert err -> ", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		log.Fatal("insert err -> ", err)
	}
	return insertId
}

func _select(sql string) (person []Person) {
	err := db.Select(&person, sql)
	if err != nil {
		log.Fatal("select err -> ", err)
	}
	return person
}

func exce(sql string, action string) {
	// update, delete
	result, err := db.Exec(sql)
	if err != nil {
		log.Fatal("exce err -> ", err)
	}
	row, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s了%d行....\n", action, row)
}

func begin() {
	t, err := db.Begin()
	updateSql := `update person set Username="s2" where user_id=10`
	result, err := t.Exec(updateSql)
	if err != nil {
		t.Rollback()
	}

	n, err := result.RowsAffected()
	if err != nil || n < 10 {
		fmt.Println(n, err, "rollback")
		t.Rollback()
	}

	// 提交之后t.close()，无法回滚
	t.Commit()
}

func main() {
	defer db.Close()
	insertSql := `insert into person(username, sex, email) values("s1", "man", "s1@qq.com")`
	insertId := insert(insertSql)
	fmt.Println("insertId -> ", insertId)

	selectSql := `select * from person where user_id=` + strconv.FormatInt(insertId, 10)
	person := _select(selectSql)
	fmt.Println("person -> ", person[0])

	updateSql := `update person set Username="s2" where user_id=` + strconv.FormatInt(insertId, 10)
	exce(updateSql, "更新")

	deleteSql := `delete from person where user_id=` + strconv.FormatInt(insertId, 10)
	exce(deleteSql, "删除")

	begin()
}
