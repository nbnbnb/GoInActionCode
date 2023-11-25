package main

import (
	"database/sql"

	_ "goinaction.zhangjin.me/chapter3/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
