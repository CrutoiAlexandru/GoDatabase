package db_commands

import (
	_ "github.com/go-sql-driver/mysql"
	"src/dbConnection"
)

func DeleteRow(val string) {
	dbc := dbConnection.Connector()

	//command to delete a row
	del, err := dbc.Query("DELETE FROM user_data WHERE user_name = '"+val+"'")

	if err != nil {
		panic(err.Error())
	}

	defer del.Close()
}

func DeleteAll() {
	dbc := dbConnection.Connector()

	//command to delete a row
	del, err := dbc.Query("DELETE FROM user_data")

	if err != nil {
		panic(err.Error())
	}

	defer del.Close()
}