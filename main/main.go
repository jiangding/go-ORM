package main

import (
	"fmt"
	orm "go-ORM"
)
import _ "github.com/mattn/go-sqlite3"

func main() {
	engine, _ := orm.NewDB("sqlite3", "haa.db")
	defer engine.Close()

	// 创建一个会话
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
