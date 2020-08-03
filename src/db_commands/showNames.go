package db_commands

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type User struct {
	name string //`json:"name"`
}

//needs conversion to a string list
func ShowNames() {
	//establish db connection
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/users")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	pickup, err := db.Query("SELECT user_name FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	for pickup.Next() {
		var user User

		err = pickup.Scan(&user.name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.name)
	}

	defer pickup.Close()
}