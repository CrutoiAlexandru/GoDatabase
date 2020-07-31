package db_commands

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteRow(val string) {
	//establish db connection
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//command to delete a row
	del, err := db.Query("DELETE FROM user_data WHERE user_name = '"+val+"'")

	if err != nil {
		panic(err.Error())
	}

	defer del.Close()
}

func DeleteAll() {
	//establish db connection
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//command to delete a row
	del, err := db.Query("DELETE FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	defer del.Close()
}