package main

import (
	"fmt"
	orm "go-ORM"
)
import _ "github.com/mattn/go-sqlite3"

func main() {
	engine, _ := orm.NewDB("sqlite3", "haa.db")
	defer engine.Close()

	fmt.Println("111")
}
