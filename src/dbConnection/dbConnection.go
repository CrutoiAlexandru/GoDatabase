package dbConnection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connector() *sql.DB{
	//establish db connection
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}

	//do I really need to close the db connection? RESEARCH
	//go db.Close()

	return db
}