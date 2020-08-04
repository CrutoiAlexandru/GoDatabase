package db_commands

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"src/dbConnection"
)

type User struct {
	name string
}

//needs conversion to a string list
func ShowNames() {
	dbc := dbConnection.Connector()

	pickup, err := dbc.Query("SELECT user_name FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	for pickup.Next() {
		var user User

		err = pickup.Scan(&user.name)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println("	"+user.name)
	}

	defer pickup.Close()
}