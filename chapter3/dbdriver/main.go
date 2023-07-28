package main

import (
	_ "GoInActionCode/chapter3/dbdriver/postgres"
	"database/sql"
)

func main() {
	sql.Open("postgres", "mydb")
}
