package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	Db, _ = sql.Open("mysql", "root:admin@/idmaker")
}

// 当前全局可用id的上界
func GetCurrentBound(category string) uint64 {
	var bound uint64
	err := Db.QueryRow("SELECT * FROM id_makers WHERE category = ? LIMIT 1", category).Scan(&category, &bound)
	checkErr(err)
	return bound
}

// 获取一组新的数（乐观锁）
func Acquire(category string, bound uint64, capacity int) bool {
	stmt, err := Db.Prepare("UPDATE id_makers SET bound = bound + ? WHERE category = ? AND bound = ?")
	checkErr(err)
	res, err := stmt.Exec(capacity, category, bound)
	checkErr(err)
	rows, _ := res.RowsAffected()
	return rows == 1
}