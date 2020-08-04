package db_commands

import (
	_ "github.com/go-sql-driver/mysql"
	"src/db_commands/idMaker"
	"src/dbConnection"
	"fmt"
)

//verify that the name does not repeat
func verifyName(val string) bool{
	dbc := dbConnection.Connector()

	pickup, err := dbc.Query("SELECT user_name FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	for pickup.Next() {
		var user string

		err = pickup.Scan(&user)

		if user == val {
			fmt.Println("Username already used!")
			return false
		}

		if err != nil {
			panic(err.Error())
		}
	}

	defer pickup.Close()

	return true
}

func AddValue(val string) {
	dbc := dbConnection.Connector()

	if verifyName(val) == true {
		//create variable and send data
		pickup, err := dbc.Query("INSERT INTO user_data (user_name) VALUES ('"+val+"')")

		if err != nil {
			panic(err.Error())
		}

		idMaker.Maker(val)
	
		defer pickup.Close()
	}
}

