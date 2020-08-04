package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/go/src/pkg/fmt"
	"github.com/jmoiron/sqlx"
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

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("mysql conn is failed err:",err)
		return
	}
	defer db.Close()
	//query := "insert into person(username,sex,email) values (?,?,?)"
	//exec, err := db.Exec(query, "baci", "nan", "12344@qq.com")
	//if err != nil {
	//	fmt.Println("Exec is failed err:",err)
	//	return
	//}
	//id, err := exec.LastInsertId()
	//if err != nil {
	//	fmt.Println("LastInsertId failed err:",err)
	//	return
	//}
	//fmt.Println(id)

	//var person []Person
	//query := "select * from person"
	//err = db.Select(&person,query, nil...)
	//if err != nil {
	//	fmt.Println("select failed err:",err)
	//	return
	//}
	//fmt.Println(person)

	//query := "update person set username=? where user_id = ?"
	//exec, err := db.Exec(query, "Tom", 2)
	//if err != nil {
	//	fmt.Println("Exec failed err:",err)
	//	return
	//}
	//id, err := exec.RowsAffected()
	//if err != nil {
	//	fmt.Println("RowsAffected failed  err:",err)
	//	return
	//}
	//fmt.Println(id)

	//query := "delete from person where user_id = ?"
	//exec, err := db.Exec(query, 2)
	//if err != nil {
	//	fmt.Println("delete failed err:",err)
	//	return
	//}
	//affected, err := exec.RowsAffected()
	//if err != nil {
	//	fmt.Println("RowsAffected failed err:",err)
	//	return
	//}
	//fmt.Println(affected)

	begin, err := db.Begin()
	if err != nil {
		fmt.Println("Begin failed err:",err)
		return
	}
	query1 := "insert into person(username,sex,email) values (?,?,?)"
	query2 := "insert into person(username,sex,email) values (?,?,?)"

	exec, err := begin.Exec(query1, "baci", "nan", "12345@qq.com")
	if err != nil {
		fmt.Println("Exec1 failed err:",err)
		begin.Rollback()
		return
	}
	_, err = exec.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed err:",err)
		begin.Rollback()
		return
	}
	exec, err = begin.Exec(query2, "baci1", "nan1", "123451@qq.com")
	if err != nil {
		fmt.Println("Exec2 failed err:",err)
		begin.Rollback()
		return
	}
	_, err = exec.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId failed err:",err)
		begin.Rollback()
		return
	}
	begin.Commit()

}