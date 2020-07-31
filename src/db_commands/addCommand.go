package db_commands

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func AddValue(val string) {
	//establish db connection
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create variable and send data
	pickup, err := db.Query("INSERT INTO user_data (user_name) VALUES ('"+val+"')")

	if err != nil {
		panic(err.Error())
	}

	defer pickup.Close()
}